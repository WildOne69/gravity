import { DnsAPIRecord, RolesDnsApi } from "gravity-api";

import { TemplateResult, html } from "lit";
import { customElement, property, state } from "lit/decorators.js";
import { ifDefined } from "lit/directives/if-defined.js";

import { DEFAULT_CONFIG } from "../../api/Config";
import "../../elements/forms/HorizontalFormElement";
import { ModelForm } from "../../elements/forms/ModelForm";

@customElement("gravity-dns-record-form")
export class DNSRecordForm extends ModelForm<DnsAPIRecord, string> {
    @property()
    zone?: string;

    @state()
    recordType = "A";

    loadInstance(pk: string): Promise<DnsAPIRecord> {
        return new RolesDnsApi(DEFAULT_CONFIG)
            .dnsGetRecords({
                zone: this.zone,
            })
            .then((records) => {
                const record = records.records?.find((z) => z.hostname + z.uid === pk);
                if (!record) throw new Error("No record");
                this.recordType = record.type;
                return record;
            });
    }

    getSuccessMessage(): string {
        if (this.instance) {
            return "Successfully updated record.";
        } else {
            return "Successfully created record.";
        }
    }

    needsRecreate(data: DnsAPIRecord): boolean {
        if (!this.instance) {
            return false;
        }
        if (data.hostname !== this.instance.hostname) return true;
        if (data.uid !== this.instance.uid) return true;
        if (data.type !== this.instance.type) return true;
        return false;
    }

    send = async (data: DnsAPIRecord): Promise<void> => {
        if (this.instance && this.needsRecreate(data)) {
            await new RolesDnsApi(DEFAULT_CONFIG).dnsDeleteRecords({
                zone: this.zone || "",
                ...this.instance,
            });
        }
        return new RolesDnsApi(DEFAULT_CONFIG).dnsPutRecords({
            zone: this.zone || "",
            ...data,
            dnsAPIRecordsPutInput: data,
        });
    };

    getLabel(): string {
        switch (this.recordType) {
            case "CNAME":
                return "CNAME Target";
            case "SRV":
                return "SRV Target";
            case "MX":
                return "Mail server";
            default:
                return "Data";
        }
    }

    renderForm(): TemplateResult {
        return html`<form class="pf-c-form pf-m-horizontal">
            <ak-form-element-horizontal label="Hostname" ?required=${true} name="hostname">
                <input
                    type="text"
                    value="${ifDefined(this.instance?.hostname)}"
                    class="pf-c-form-control"
                    required
                />
            </ak-form-element-horizontal>
            <ak-form-element-horizontal label="UID" ?required=${true} name="uid">
                <input
                    type="number"
                    value="${this.instance?.uid || 0}"
                    class="pf-c-form-control"
                    required
                />
                <p class="pf-c-form__helper-text">
                    Unique identifier to configure multiple records for the same hostname.
                </p>
            </ak-form-element-horizontal>
            <ak-form-element-horizontal label="Type" ?required=${true} name="type">
                <select
                    class="pf-c-form-control"
                    @change=${(ev: Event) => {
                        const current = (ev.target as HTMLInputElement).value;
                        this.recordType = current;
                    }}
                >
                    <option value="A" ?selected=${this.instance?.type === "A"}>A</option>
                    <option value="AAAA" ?selected=${this.instance?.type === "AAAA"}>AAAA</option>
                    <option value="CNAME" ?selected=${this.instance?.type === "CNAME"}>
                        CNAME
                    </option>
                    <option value="PTR" ?selected=${this.instance?.type === "PTR"}>PTR</option>
                    <option value="NS" ?selected=${this.instance?.type === "NS"}>NS</option>
                    <option value="MX" ?selected=${this.instance?.type === "MX"}>MX</option>
                    <option value="SRV" ?selected=${this.instance?.type === "SRV"}>SRV</option>
                </select>
            </ak-form-element-horizontal>
            <ak-form-element-horizontal label="${this.getLabel()}" ?required=${true} name="data">
                <input
                    type="text"
                    value="${ifDefined(this.instance?.data)}"
                    class="pf-c-form-control"
                    required
                />
            </ak-form-element-horizontal>
            ${this.recordType === "MX"
                ? html`
                      <ak-form-element-horizontal
                          label="MX Preference"
                          ?required=${true}
                          name="mxPreference"
                      >
                          <input
                              type="number"
                              value="${ifDefined(this.instance?.mxPreference)}"
                              class="pf-c-form-control"
                              required
                          />
                      </ak-form-element-horizontal>
                  `
                : html``}
            ${this.recordType === "SRV"
                ? html`
                      <ak-form-element-horizontal label="SRV Port" ?required=${true} name="srvPort">
                          <input
                              type="number"
                              value="${ifDefined(this.instance?.srvPort)}"
                              class="pf-c-form-control"
                              required
                          />
                      </ak-form-element-horizontal>
                      <ak-form-element-horizontal
                          label="SRV Priority"
                          ?required=${true}
                          name="srvPriority"
                      >
                          <input
                              type="number"
                              value="${ifDefined(this.instance?.srvPriority)}"
                              class="pf-c-form-control"
                              required
                          />
                      </ak-form-element-horizontal>
                      <ak-form-element-horizontal
                          label="SRV Weight"
                          ?required=${true}
                          name="srvWeight"
                      >
                          <input
                              type="number"
                              value="${ifDefined(this.instance?.srvWeight)}"
                              class="pf-c-form-control"
                              required
                          />
                      </ak-form-element-horizontal>
                  `
                : html``}
        </form>`;
    }
}
