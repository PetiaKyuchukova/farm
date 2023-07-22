import { customElement, property, state, query} from 'lit/decorators.js'
import { LitElement, html, css } from 'lit'
import {Task} from "./task.type.ts";
import "../cow-profile/profile.ts"

@customElement('farm-tasks')
export class FarmTasks extends LitElement {
    static styles = css`
 
    .content {
    background: white;
    width: 80%;
    left: 10%;
    top: 20%;
    position: absolute;
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


    FertilizationType     = "AI"
    PregnantType          = "PC"
    DryPeriodAfter15dType = "DP15d"
    DryPeriodStartType    = "DPS"
    GivingBirthType       = "GB"
    OvulationType         = "OVU"
    PostMilkType         = "M"

    warning = "table-warning"
    success = "table-success"
    danger = "table-danger"
    info = "table-info"

    mapTaksTypes = new Map();

    fillMap(){
       let  mapTypes = new Map();
        mapTypes.set(this.FertilizationType, this.warning)
        mapTypes.set(this.PregnantType, this.warning)

        mapTypes.set(this.GivingBirthType, this.danger)
        mapTypes.set(this.DryPeriodStartType, this.danger)

        mapTypes.set(this.OvulationType, this.success)
        mapTypes.set(this.DryPeriodAfter15dType, this.success)
        mapTypes.set(this.PostMilkType, this.info)

        return mapTypes
    }

    firstUpdated() {
        this.fetchData()
        this.mapTaksTypes = this.fillMap()
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


        if (this.error !== ''){
            return html`
            <div style="    
            background: white;
            top: 50%;
            left: 40%;
            padding: 20px;
            position: absolute;
            border-radius: 10px;"> Ooops...something get wrong!</div>
            `
        }
        let rows= []

        if (this.data.length > 0){
            for (const task of this.data) {
                let row = html`
                    <tr class="${this.mapTaksTypes.get(task.type)}" @click=${this.openCowProfile(task.cow_id)}>
                        <td>${task.cow_id}</td>
                        <td>${task.text}</td>
                    </tr>`
                rows.push(row)
            }
        }

        return html`
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
            <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>

            
            <div class="content">
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