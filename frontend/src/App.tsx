import React, { useState } from 'react';
import './Main.css'
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { redirect } from 'react-router-dom';
function App() {
  const [login, setLogin] = useState('')
  const [password, setPassord] = useState('')
  const makeLogin = () => {
    if (!login) {
      toast.error("Login não pode ser vazio", { toastId: "error" })
      return
    }
    if (!password) {
      toast.error("Senha não pode ser vazia", { toastId: "error" })
      return
    }
    const body = JSON.stringify({ login: login, password: password })
    let headers = new Headers()
    headers.append("Content-Type", "application/json")
    fetch("/login", { method: "post", body: body, headers: headers }).then(async (res) => {
      let msg = await res.json()
      if (res.status !== 200) {
        toast.error(msg.error, { toastId: "error" })
        return
      }
      localStorage.setItem("token", msg.token)
      window.location.pathname = "/users"
    }).catch(err => toast.error(err, { toastId: "error" }))
  }
  return (
    <div className="relative flex flex-col justify-center min-h-screen overflow-hidden">
      <div className="w-full p-6 m-auto bg-white rounded-md shadow-md lg:max-w-xl">
        <h1 className="text-3xl font-semibold text-center text-purple-700 underline">
          Sign in
        </h1>
        {/* <form className="mt-6"> */}
        <div className="mb-2">
          <label
            className="block text-sm font-semibold text-gray-800"
          >
            Login
          </label>
          <input
            type="text"
            className="block w-full px-4 py-2 mt-2 text-white bg-purple-700  border rounded-md focus:border-purple-400 focus:ring-purple-300 focus:outline-none focus:ring focus:ring-opacity-40"

            value={login}
            onChange={e => setLogin(e.target.value)}

          />
        </div>
        <div className="mb-2">
          <label

            className="block text-sm font-semibold text-gray-800"
          >
            Password
          </label>
          <input
            type="password"
            className="block w-full px-4 py-2 mt-2 text-white bg-purple-700 border rounded-md focus:border-purple-400 focus:ring-purple-300 focus:outline-none focus:ring focus:ring-opacity-40"
            onChange={e => setPassord(e.target.value)}
            value={password}
          />
        </div>
        <div className='mt-6'>
          <button onClick={makeLogin} className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200  bg-purple-700 rounded-md hover:bg-purple-600 focus:outline-none focus:bg-purple-600">Login</button>
        </div>
      </div>
      <ToastContainer />
    </div>
  )
}

export default App;
