import { customElement, property, query} from 'lit/decorators.js'
import {LitElement, html, css} from 'lit'
import {Milk} from "./milk.type.ts";
import '../milk-post/widget.ts'


@customElement('farm-milk-list')
export class FarmMilkList extends LitElement {
    static styles = css`
 
   .content {
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
    data: Milk[] = []

    @property({attribute: false, type: Boolean})
    isLoading = false

    @query("#postMilk")
    postMilk: HTMLElement

    totalProfit = 0

    private fetchData() {
        const date = new Date();
        const firstDayCurrentMonth = this.getFirstDayOfMonth(
            date.getFullYear(),
            date.getMonth(),
        );

        const lastDayCurrentMonth = this.getLastDayOfMonth(
            date.getFullYear(),
            date.getMonth(),
        );

        this.updateComplete.then(() => {
            this.isLoading = true

            fetch(`http://localhost:9030/milk?from=${firstDayCurrentMonth}&to=${lastDayCurrentMonth}`)
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
     getFirstDayOfMonth(year: number, month:number) {
         var d = new Date(year, month, 1),
             monthN = '' + (d.getMonth() + 1),
             dayN = '' + d.getDate(),
             yearN = d.getFullYear();

         if (monthN.length < 2)
             monthN = '0' + monthN;
         if (dayN.length < 2)
             dayN = '0' + dayN;

         return [yearN, monthN, dayN].join('-');
    }

     getLastDayOfMonth(year: number, month:number) {
         var d = new Date(year, month + 1, 0),
             monthN = '' + (d.getMonth() + 1),
             dayN = '' + d.getDate(),
             yearN = d.getFullYear();

         if (monthN.length < 2)
             monthN = '0' + monthN;
         if (dayN.length < 2)
             dayN = '0' + dayN;

         return [yearN, monthN, dayN].join('-');
    }


    firstUpdated() {
        this.fetchData()
    }

    private openPostForm() {
        return (_e: MouseEvent) => {
            this.visible = !this.visible
            this.postMilk.setAttribute("visible", "up")
        }
    }

   render(){
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

       let rows = []
       this.totalProfit =0

       if (this.data!=undefined){
           for (const milk of this.data) {
               this.totalProfit += milk.liters * milk.price

               let row = html`
                <tr>
                    <td>${milk.date}</td>
                    <td>${milk.liters}</td>
                    <td>${milk.price}</td>
                    <td class="table-active" >${milk.liters * milk.price}</td>
                </tr>
                `
               rows.push(row)
           }
       }


       return html`
           <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM" crossorigin="anonymous">
           <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js" integrity="sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz" crossorigin="anonymous"></script>


           <div class="content">
               <div style="display: flex; justify-content: space-between;">
                   <h1>Monthly Milk</h1>
                   <button type="button" style="height: 40px;" class="btn btn-success" @click=${this.openPostForm()}>+ Add Milk</button>
               </div>
               <table class="table">
                   <thead>
                   <th>Date</th>
                   <th>Liters</th>
                   <th>PRICE/liter</th>
                   <th >Price</th>
                   </thead>
                   <tbody>
                   ${rows}
                   </tbody>
               </table>
               <div style="text-align: right"> 
                <h5>Total profit: </h5>
                <p>${this.totalProfit} </p>
               </div>
           </div>
           
           <farm-milk-post id="postMilk"></farm-milk-post>
       `
   }


}