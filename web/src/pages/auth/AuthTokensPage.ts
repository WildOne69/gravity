import { AuthToken, RolesApiApi } from "gravity-api";

import { TemplateResult, html } from "lit";
import { customElement } from "lit/decorators.js";

import { DEFAULT_CONFIG } from "../../api/Config";
import "../../elements/forms/DeleteBulkForm";
import "../../elements/forms/ModalForm";
import { PaginatedResponse, TableColumn } from "../../elements/table/Table";
import { TablePage } from "../../elements/table/TablePage";
import { PaginationWrapper } from "../../utils";
import "./AuthTokenForm";

@customElement("gravity-auth-tokens")
export class AuthTokensPage extends TablePage<AuthToken> {
    pageTitle(): string {
        return "Tokens";
    }
    pageDescription(): string | undefined {
        return undefined;
    }
    pageIcon(): string {
        return "";
    }
    checkbox = true;

    searchEnabled(): boolean {
        return true;
    }

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    apiEndpoint(page: number): Promise<PaginatedResponse<AuthToken>> {
        return new RolesApiApi(DEFAULT_CONFIG).apiGetTokens().then((tokens) => {
            const data = (tokens.tokens || []).filter((l) =>
                l.username.toLowerCase().includes(this.search.toLowerCase()),
            );
            data.sort((a, b) => {
                if (a.username > b.username) return 1;
                if (a.username < b.username) return -1;
                return 0;
            });
            return PaginationWrapper(data);
        });
    }

    columns(): TableColumn[] {
        return [new TableColumn("Username")];
    }

    row(item: AuthToken): TemplateResult<1 | 2>[] {
        return [html`${item.username}`];
    }

    renderToolbarSelected(): TemplateResult {
        const disabled = this.selectedElements.length < 1;
        return html`<ak-forms-delete-bulk
            objectLabel=${"Tokens(s)"}
            .objects=${this.selectedElements}
            .metadata=${(item: AuthToken) => {
                return [{ key: "Username", value: item.username }];
            }}
            .delete=${(item: AuthToken) => {
                return new RolesApiApi(DEFAULT_CONFIG).apiDeleteTokens({
                    key: item.key,
                });
            }}
        >
            <button ?disabled=${disabled} slot="trigger" class="pf-c-button pf-m-danger">
                ${"Delete"}
            </button>
        </ak-forms-delete-bulk>`;
    }

    renderObjectCreate(): TemplateResult {
        return html`
            <ak-forms-modal>
                <span slot="submit"> ${"Create"} </span>
                <span slot="header"> ${"Create Token"} </span>
                <gravity-auth-token-form slot="form"> </gravity-auth-token-form>
                <button slot="trigger" class="pf-c-button pf-m-primary">${"Create"}</button>
            </ak-forms-modal>
        `;
    }
}