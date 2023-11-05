import './Main.css'
import Table, { Irow, Icolumn } from 'react-tailwind-table'
import "react-tailwind-table/dist/index.css"; //use for a non tailwind project
import User from './types/User';

type Data = {
    users: User[]
}

const Users = ({ users }: Data) => {
    let col = [
        {
            field: "name",
            use: "Nome"
        },
        {
            // use_in_display: false,
            field: "email", //Object destructure
            use: "Email"
        },
        {
            field: "validate",
            use: "Situacão"
        },
        {
            field: "formated_date",
            use: "Data"
        },
        {
            field: "op",
            use: "Opcão"
        },
        {
            field: "id",
            use_in_display: false
        }
    ];
    let rowChecker = (row: Irow, col: Icolumn, display_value: string): JSX.Element | string => {
        if (col.field === "validate") {
            return <p className={row[col.field] ? "success" : "error"}>{row[col.field] ? "Pago" : "Atrasado"}</p>
        }
        if (col.field === "op") {
            let id = row["id"]
            let url = `/edit/${id}`
            return <a className="btn" href={url}>Editar</a>
        }

        return display_value
    }
    let on_search = (search_word: string, result?: Irow[] | []): void => {

    }
    return (
        <>
            <Table columns={col} rows={users} per_page={5} row_render={rowChecker} should_export={false} on_search={on_search} />
            <a className='add-user-btn btn' href='/add'>
                Adicionar
            </ a>
        </>
    )

}





export default Users