import { useState } from "react";
import User from "./types/User";
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';
type Data = {
    user: User
}
const Edit = ({ user }: Data) => {
    const [name, setName] = useState(user.name)
    const [email, setEmail] = useState(user.email)
    const [date, setDate] = useState("")
    const updateUser = () => {
        if (!name) {
            toast.error("Nome vazio amigao?", { toastId: "error" })
            return
        }
        if (!date) {
            toast.error("Data vazia amigao?", { toastId: "error" })
            return
        }
        let s = date.split("-")
        let d = new Date(parseInt(s[0]), parseInt(s[1]), parseInt(s[2]))
        let year = d.getFullYear()
        let mounth = d.getMonth()
        let day: number | string = d.getDate()
        if (day < 10) {
            day = '0' + day
        }
        let format = `${year}-${mounth}-${day}`
        let body = JSON.stringify({ id: user.id, email: email, formated_date: format, name: name })
        let headers = new Headers()
        headers.append("Content-Type", "application/json")
        fetch("/edit", { method: "put", body: body, headers: headers }).then(async (res) => {
            let msg = await res.json()
            if (res.status !== 200) {
                toast.error(msg.error, { toastId: "error" })
                return
            }
            window.location.pathname = "/users"
        }).catch(err => toast.error(err, { toastId: "error" }))
    }
    const deleteUser = () => {
        let headers = new Headers()
        headers.append("Content-Type", "application/json")
        fetch(`/delete/${user.id}`, { headers: headers, method: "delete" }).then(async (res) => {
            if (res.status != 200) {
                let msg = await res.json()
                toast.error(msg.err)
                return
            }
            window.location.pathname = "/users"
        })
    }
    return <>
        <div className="relative flex flex-col justify-center min-h-screen overflow-hidden">
            <div className="w-full p-6 m-auto bg-white rounded-md shadow-md lg:max-w-xl">
                <h1 className="text-3xl font-semibold text-center text-purple-700 underline">
                    Usuario
                </h1>
                {/* <form className="mt-6"> */}
                <div className="mb-2">
                    <label className="block text-sm font-semibold text-gray-800">
                        Nome
                    </label>
                    <input type="text" className="block w-full px-4 py-2 mt-2 text-white bg-purple-700  border rounded-md focus:border-purple-400 focus:ring-purple-300 focus:outline-none focus:ring focus:ring-opacity-40"
                        value={name}
                        onChange={e => setName(e.target.value)}

                    />
                </div>
                <div className="mb-2">
                    <label className="block text-sm font-semibold text-gray-800">
                        Email
                    </label>
                    <input
                        type="email"
                        className="block w-full px-4 py-2 mt-2 text-white bg-purple-700 border rounded-md focus:border-purple-400 focus:ring-purple-300 focus:outline-none focus:ring focus:ring-opacity-40"
                        onChange={e => setEmail(e.target.value)}
                        value={email}
                    />
                </div>

                <div className="mb-2">
                    <label className="block text-sm font-semibold text-gray-800">
                        Data
                    </label>
                    <input type="date" className="block w-full px-4 py-2 mt-2 text-white bg-purple-700  border rounded-md focus:border-purple-400 focus:ring-purple-300 focus:outline-none focus:ring focus:ring-opacity-40"
                        onChange={e => {
                            setDate(e.target.value)
                        }}
                    />
                </div>
                <div className='mt-6'>
                    <button onClick={updateUser} className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200  bg-purple-700 rounded-md hover:bg-purple-600 focus:outline-none focus:bg-purple-600">Salvar</button>
                </div>
                <div className='mt-6'>
                    <button onClick={deleteUser} className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200  bg-purple-700 rounded-md hover:bg-purple-600 focus:outline-none focus:bg-purple-600">Deletar</button>
                </div>
            </div>
            <ToastContainer />
        </div>
    </>
}

export default Edit;