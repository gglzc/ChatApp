import React, {useState , useEffect ,useContext} from 'react'
import { API_URL, WEBSOCKET_URL } from '../constants'
import { v4 as uuidv4} from 'uuid'
import { useRouter } from 'next/router'
import { AuthContext, UserInfo } from '../module/auth_provider'
import { WebsocketContext } from '../module/websocket_provider'



const index = () => {
  const [rooms ,setRooms] = useState<{id: string , name:string}[]>([])
  const [roomName , setRoomName] = useState('')
  const { user } = useContext(AuthContext)
  const { setConn } =useContext(WebsocketContext)
  const router =useRouter()
  
 
  
const getRooms = async() =>{
    try{
      const res = await fetch(`${API_URL}/ws/getRooms` ,{
        method: 'GET',
      })

      const data =await res.json()
      if (res.ok){
        setRooms(data)
      }
    } catch(err){
      console.log(err)
    }
  }
  
  useEffect(() =>{
    getRooms()
  },[])
  
  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault()

    try{
      setRoomName('')
      const res = await fetch( `${API_URL}/ws/createroom` , {
        method:'POST',
        headers: { 'Content-Type' : 'application/json'},
        credentials: 'include',
        body: JSON.stringify({ 
          id : uuidv4(),
          name: roomName,
        }),
     })
     if(res.ok){
        getRooms()
     }
    }catch(err){
      console.log(err)
     }
}

const logoutSubmitHandler =async(e:React.SyntheticEvent)=>{
  e.preventDefault()

  
  try {
    const res =await fetch(`${API_URL}/logout`,{
      method:'GET',
      headers: { 'Content-Type' : 'application/json'},
      credentials: 'include',
    })

    
    if (res.ok){
      localStorage.clear()  
      router.push('/login')
    }
  } catch (error) {
    console.log(error)
  }
}


const joinRoom = (roomId: string ) =>{
  const ws = new WebSocket(`${WEBSOCKET_URL}/ws/joinRoom/${roomId}?userId=${user.id}&username=${user.username}`)
    if (ws.OPEN){
      setConn(ws)
      router.push('/app')
      return
    }
}

  
  return (
   <>
      <div className='my-8 px-4 md:mx-32 w-full h-full'>
        <div className='flex justify-center mt-3 p-5'>
          <input 
          type='text' 
          className='border border-blue p-2 rounded-md focus:outline-none focus:border-red' 
          placeholder='ROOM NAME'
          value={roomName}
          onChange={(e) => setRoomName(e.target.value)}
          />
          <button 
          className='border border-blue p-2 rounded-md  md:ml-4' 
          onClick={submitHandler}>
            建立房間
          </button>
          <button
          className='border border-blue p-2 rounded-md  md:ml-6'
          onClick={logoutSubmitHandler}>
            登出
          </button>
        </div>
        <div className='mt-6'>
          <div className='font-bold'>聊天室</div>
          <div className='grid grid-cols-1 md:grid-cols-5 gap-4 mt-6'>
            {rooms.map((room, index) => (
              <div
                key={index}
                className='border border-blue p-4 flex items-center rounded-md w-full'
              >
                <div className='w-full'>
                  <div className='text-sm'>房間</div>
                  <div className='text-blue font-bold text-lg'>{room.name}</div>
                </div>
                <div className=''>
                  <button
                    className='px-4 text-white bg-blue rounded-md'
                    onClick={() => joinRoom(room.id)}
                  >
                    加入房間
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  )
}


export default index