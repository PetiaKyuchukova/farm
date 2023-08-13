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

    private updateTask(task: Task) {
        console.log(task)
        fetch(`http://localhost:9030/tasks/update`, {
            method: 'PUT',
            body: JSON.stringify(task)
        }).then(async (response) => {
            if (response.ok) {
                console.log("Saved!")
            } else {
                this.error = 'Error updating task.'
                console.log("err saving cow!")
            }
        })

        window.location.reload();
        return
    }
    private openCowProfile(cowId: string) {
        return (_e: MouseEvent) => {
            this.idx = cowId
            this.visible = !this.visible
        }
    }

    private taskStatus(taskId: number) {
        return (_e: MouseEvent) => {
           this.data[taskId].done = !this.data[taskId].done
            console.log("task",this.data[taskId])

            this.updateTask(this.data[taskId])
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
        mapTypes.set(this.PregnantType, this.warning)


        mapTypes.set(this.GivingBirthType, this.danger)
        mapTypes.set(this.DryPeriodStartType, this.danger)
        mapTypes.set(this.OvulationType, this.danger)
        mapTypes.set(this.FertilizationType, this.danger)

        mapTypes.set(this.DryPeriodAfter15dType, this.success)
        mapTypes.set(this.PostMilkType, this.info)

        return mapTypes
    }

    firstUpdated() {
        this.fetchData()
        this.mapTaksTypes = this.fillMap()
    }

    render() {

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
            for (let i = 0; i < this.data.length; i++) {
                let row = html`
                    <tr style="text-decoration: ${this.data[i].done ? "line-through" : "none"}; opacity: ${this.data[i].done ? 0.5 : 1}" class="${this.mapTaksTypes.get(this.data[i].type)}"  >
                        <td style="cursor: pointer" @click=${this.openCowProfile(this.data[i].cow_id)}>${this.data[i].cow_id}</td>
                        <td>${this.data[i].text}</td>
                        <td style="text-align: right; background: gainsboro;" @click=${this.taskStatus(i)}><i class="bi bi-check-lg"></i></td>
                    </tr>`
                rows.push(row)
            }
        }

        return html`
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
            <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>
            <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css">
            
            <div class="content">
                <h1>Tasks for today </h1>
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