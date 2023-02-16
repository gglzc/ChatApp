import { useRouter } from 'next/router'
import React, { useState  , useContext , useEffect} from 'react'
import { API_URL } from '../../constants'
import { AuthContext,UserInfo } from '../../module/auth_provider'

const index = () => {
const [email , setEmail] = useState ('')
const [password , setPassword] = useState('')
const { auth } = useContext(AuthContext)

const router = useRouter()

//login and share information  across with different pages
useEffect(() => {
    if(auth){
        router.push('/')
        return
    }
}, [auth])


//handle submit function logic
const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault()

    try{
     const res = await fetch( `${API_URL}/login` , {
        method:'POST',
        headers: { 'Content-Type' : 'application/json'},
        body: JSON.stringify({ email,password })
     })

     const data  = await res.json()
     if(res.ok){
        const user:UserInfo ={
            username : data.username,
            id: data.id,
        }

        localStorage.setItem('user_info' , JSON.stringify(user))
        return router.push('/')
     }
    }catch(err) {
        console.log(err)
    }
}
//
const SignupSubmitHandler = async (e:React.SyntheticEvent)=>{
    e.preventDefault()
    return router.push('/signup')
}


return (
    <div className='flex items-center justify-center min-w-full min-h-screen'>
        <form className='flex flex-col md:w-1/5'>
            <div className='text-3xl font-bold text-center border-gray'>
                <span className='text-red'>OMA Chat</span>
            </div>
            <input 
            placeholder='email' 
            className='p-3 mt-8 rounded-md border-2 border-b-orange-500 focus:outline-none focus:border-orange-800'
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            />
            <input 
            type='password' 
            placeholder='password' 
            className='p-3 mt-4 rounded-md border-2 border-b-orange-500 focus:outline-none focus:border-orange-800'
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            />
            <button className='p-3 mt-6 rounded-sm bg-red font-bold text-white' onClick={submitHandler}>
            login
            </button>
            
            <div className="flex justify-center mt-6">
                <button className="p-3 mt-6 rounded-md bg-blue font-bold text-white" onClick={SignupSubmitHandler}>
                申請帳號
                </button>
            </div>
        </form>
    </div>
  )
}

export default index