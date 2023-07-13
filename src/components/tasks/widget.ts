import { customElement, property, state, query} from 'lit/decorators.js'
import { LitElement, html, nothing, css , hidden} from 'lit'
import {Task} from "./task.type.ts";
import "../cow-profile/profile.ts"

@customElement('farm-tasks')
export class FarmTasks extends LitElement {
    static styles = css`
 
    .content {
    background: white;
    max-width: 80%;
    margin-left: 10%;
    margin-top: 20%;
    border-radius: 24px;
    padding: 20px;
}
h1 {
color: #367749
}
    `


    @property({attribute: false, type: String})
    error = ''

    @property({attribute: false, type: Array})
    data: Task[] = []

    @state()
    private visible = false

    @state()
    private idx = ''

    @property({attribute: false, type: Boolean})
    isLoading = false

    @query('#profile')
    private profile: HTMLElement

    private fetchData() {
        this.updateComplete.then(() => {
            this.isLoading = true

            fetch(`http://localhost:9030/tasks?date=2023-07-02`)
                .then(async resp => {
                    console.log(this.data)
                    if (resp.status === 200) {
                        this.data = await resp.json()
                    } else {
                        this.error = 'Error loading cost anomalies.'
                    }
                    this.isLoading = false
                })
                .catch(reason => {
                    this.error = `Network error: ${reason}`
                    this.isLoading = false
                })
        })
    }

    private openCowProfile(cowId: string) {
        return (_e: MouseEvent) => {
            this.idx = cowId
            this.visible = !this.visible
        }
    }

    private closeCowProfile(cowId: string) {
        return (_e: MouseEvent) => {
            this.idx = cowId
            this.visible = true
        }
    }


    firstUpdated() {
        this.fetchData()
    }
    render() {
        if (this.isLoading) {
            return html`
                <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
                <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>

                <div class="spinner-border" role="status">
                    <span class="visually-hidden">Loading...</span>
                </div>`
        }
        let rows= []

        if (this.data.length > 0){
            for (const task of this.data) {
                let row = html`<tr @click=${this.openCowProfile(task.cow_id)}>
                <td>${task.cow_id}</td>
                <td>${task.text}</td>
            </tr>`
                rows.push(row)
            }
        }

        return html`
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
            <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>

            
            <div class="content" style="${this.visible? 'filter: blur(8px);-webkit-filter: blur(8px);' : ''}">
                <h1>Tasks</h1>
                <table class="table table-hover">
                    <thead>
                    <td>Cow Id</td>
                    <td>Task</td>
                    </thead>
                    <tbody>
                    ${rows}
                    </tbody>
                </table>
            </div>
            <farm-cow-profile  id="profile" cow="${this.idx}" visible="${this.visible}"></farm-cow-profile>
        `}

}