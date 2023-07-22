import { customElement, property,query,state} from 'lit/decorators.js'
import {LitElement, html, css,} from 'lit'
import {Cow} from "./cow.type.ts";
import "../cow-profile/profile.ts"

@customElement('farm-herd')
export class FarmHerd extends LitElement {
    static styles = css`
 
    .content {
        background: white;
        width: 80%;
        position: absolute;
        left: 10%;
        top: 30%;
        border-radius: 24px;
        padding: 20px;
    }
    .search {
        background: white;
        width: 80%;
        position: absolute;
        left: 10%;
        top: 20%;
        border-radius: 24px;
        padding: 20px;
    }
    h1 {
        color: #367749
    }
    `
    @property({attribute: false, type: String})
    error = ''

    @property({attribute: true, type: Boolean})
    visible = false

    @property({attribute: false, type: Array})
    data: Cow[] = []

    @property({attribute: false, type: Boolean})
    isLoading = false

    @query("#myInput")
    myInput: HTMLInputElement

    @query("#myTable")
    myTable: HTMLElement

    @state()
    private idx = ''


    private fetchData() {
        this.updateComplete.then(() => {
            this.isLoading = true

            fetch(`http://localhost:9030/cows`)
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
            if (cowId !== ""){
                this.idx = cowId
            }
            console.log('click')
            this.visible = !this.visible
            console.log(this.visible)

        }
    }

    firstUpdated() {
        this.fetchData()
    }

    private search(e: any) {

        var filter, table, tr, td, i,j, txtValue;
        filter = e.target.value.toUpperCase();
        table = this.myTable
        tr = table.getElementsByTagName("tr");

        for (i = 0; i < tr.length; i++) {
            td = tr[i].getElementsByTagName("td");
            for (j = 0; j < td.length; j++) {
                if (td[j]) {
                    txtValue = td[j].textContent || td[j].innerText;
                    if (txtValue.toUpperCase().indexOf(filter) > -1) {
                        tr[i].style.display = "";
                        break
                    } else {
                        tr[i].style.display = "none";
                    }
                }
            }
        }
    }
    private renderTable(){
        let rows = []
        if (this.data!=undefined){
            for (const cow of this.data) {
                let row = html`
                <tr @click=${this.openCowProfile(cow.id)}>
                    <td>${cow.id}</td>
                    <td>${cow.gender}</td>
                    <td>${cow.breed}</td>
                    <td>${cow.colour}</td>
                </tr>
                `
                rows.push(row)
            }
        }

        return rows
    }

    render() {

        let rows = this.renderTable()

        return html`
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
            <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>
            
            <div class="search">
                <input type="text" id="myInput" @keyup="${this.search}" placeholder="Search for cows.." title="Type in a name">
            </div>
            <div class="content">
                <div style="display: flex;     justify-content: space-between;">
                    <h1>Herd</h1>
                    <button type="button" style="height: 40px;" class="btn btn-success" @click=${this.openCowProfile("new")}>+ Add cow</button>
                    
                </div>
                <table id="myTable" class="table table-hover">
                    <thead>
                    <tr>
                        <th>Cow ID</th>
                        <th>Gender</th>
                        <th>Breed</th>
                        <th>Color</th>
                    </tr>
                    </thead>
                    <tbody>
                        ${rows}
                    </tbody>
                </table>
            </div>

            <farm-cow-profile  id="profile" cow="${this.idx}" visible="${this.visible}"></farm-cow-profile>
            
           
        `}
}

