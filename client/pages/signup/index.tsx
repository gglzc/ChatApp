import { useRouter } from 'next/router'
import React from 'react'
import { useState , SyntheticEvent} from 'react'
import {API_URL} from '../../constants/index'

const index = () => {
    const [username , setUsername] = useState('')
    const [email , setEmail] = useState('')
    const [password ,setPassword] = useState('')
    const [confirm_password ,setConfirmPassword] = useState('')
    const router = useRouter()
  
    const submitHandler = async (e:SyntheticEvent) =>{
      e.preventDefault()
      //先check密碼正不正確
      if (password !== confirm_password) {
        setPassword('')
        setConfirmPassword('')
        alert('密碼不一致，請重新確認密碼')
        return
      }
      //interact with backend
      await fetch( `${API_URL}/signup`,{
          method: 'POST',
          headers: {'Content-Type':'application/json'},
          body: JSON.stringify({
            username,
            password,
            email,
          }),
      })
        //申請完轉換頁面到login page 若申請失敗則留在原畫面並且告知是什麼原因
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            alert(response.error);
          }
          //沒有錯誤並且查看有沒有response ok才能跳轉
          if (!response.error) {
            if(response.ok){
            return router.push('/login');
            }
          }
        });
    }

return (
<>
<div className='flex items-center justify-center min-w-full min-h-screen'>
  <form className='flex flex-col md:w-1/5'>
      <div className='text-3xl font-bold text-center border-gray'>
          <span className='text-red'>申請帳號</span>
      </div>
        <input 
        placeholder='email' 
        className='p-3 mt-8 rounded-md border-2 border-b-orange-500 focus:outline-none focus:border-orange-800'
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        />
        <input 
        placeholder='Username' 
        className='p-3 mt-4 rounded-md border-2 border-b-orange-500 focus:outline-none focus:border-orange-800'
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        />
        <input 
        type='password' 
        placeholder='password' 
        className='p-3 mt-4 rounded-md border-2 border-b-orange-500 focus:outline-none focus:border-orange-800'
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        />
        <input 
        type='password' 
        placeholder='Input password again' 
        className='p-3 mt-4 rounded-md border-2 border-b-orange-500 focus:outline-none focus:border-orange-800'
        value={confirm_password}
        onChange={(e) => setConfirmPassword(e.target.value)}
        />
        <button 
        className='p-3 mt-6 rounded-sm bg-red font-bold text-white' 
        onClick={submitHandler}>
        送出
        </button>
  </form>
</div>
</>
)
}

export default index