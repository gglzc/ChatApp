import { useState ,createContext , useEffect } from 'react'
import { useRouter } from 'next/router'


export type UserInfo = {
    username : string
    id : string
}

export  const AuthContext = createContext<{
    auth:boolean
    setAuth: (auth: boolean) => void 
    user: UserInfo,
    setUser:(user: UserInfo) =>void
}>({
    auth:false,
    setAuth: () => {},
    user:{username: '' , id:''},
    setUser:() => {},
})


const AuthContextProvider = ({children} : {children: React.ReactNode}) => {
    const [auth,setAuth]= useState(false)
    const [user ,setUser] =useState<UserInfo>({username:'' , id:''})
  
    const router = useRouter()
    
    useEffect(() => {
        //先拿user的info
        const userInfo = localStorage.getItem('user_info')
        //如果沒有則跳轉到申請畫面
        if (!userInfo){
            if(window.location.pathname != '/singup'){
                router.push('/login')
                return
            }
        }else{ //有資料 則開始渲染更新畫面及資料
            const user : UserInfo = JSON.parse(userInfo)
                if (user){
                    setUser({
                        username: user.username,
                        id: user.id,
                    })
                }
            //授權auth成功
            setAuth(true)
        }
    }, [auth])
    
    return (
     <> 
      <AuthContext.Provider 
        value={{
            auth:auth,
            setAuth:setAuth,
            user: user,
            setUser: setUser,
        }}
        >
            {children}
      </AuthContext.Provider>
     </>
    )
}

export default AuthContextProvider