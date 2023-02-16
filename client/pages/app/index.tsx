import { useState , useEffect ,useContext ,useRef} from "react"
import React from 'react'
import ChatBody from "../../components/chat_body"
import { WebsocketContext } from "../../module/websocket_provider"
import { useRouter } from "next/router"
import { API_URL } from "../../constants"
import { AuthContext } from "../../module/auth_provider"




export type Message = {
    content: string
    client_id: string
    username: string
    room_id: string
    type: 'recv' | 'self' //收到訊息或是自己發的
  }

  const index = () => {
    const [messages , setMessages] = useState<Array<Message>>([])
    const textarea = useRef<HTMLTextAreaElement>(null)
    const {conn} = useContext(WebsocketContext)
    const router = useRouter()
    const [users , setUsers] = useState<Array< { username: string }>>([])
    const {user} = useContext(AuthContext)
    const {auth} = useContext(AuthContext)
        
    
    //get clients in the room
        useEffect(() => {
            if (conn === null) {
              router.push('/')
              return
            }
        
            const roomId = conn.url.split('/')[5]
            async function getUsers() {
              try {
                const res = await fetch(`${API_URL}/ws/getClients/${roomId}`, {
                  method: 'GET',
                  headers: { 'Content-Type': 'application/json' },
                })
                const data = await res.json()
        
                setUsers(data)
              } catch (e) {
                console.error(e)
              }
            }
            getUsers()
          }, [])
          //handle websocket connection
          useEffect(() => {
            if (textarea.current) {
                textarea.current
            }
        
            if (conn === null) {
              router.push('/')
              return
            }

            conn.onmessage =(message) =>{
                const m:Message = JSON.parse(message.data)
                if (m.content == 'A new User Come the room'){
                    setUsers([...users , {username:m.username}])
                }
                if (m.content == 'user left the chat'){
                    const deleteUser = users.filter((user)=> user.username != m.username)
                    setUsers([...deleteUser])
                    setMessages([...messages , m])
                    return
                }

                user?.username == m.username ? (m.type= 'self') : (m.type='recv')
                setMessages([...messages , m])
            }

            conn.onclose = () =>{}
            conn.onerror =() =>{}
            conn.onopen = ()=>{}
          },[textarea,messages,conn,users])
    
    const sendMessage=()=>{
        if (!textarea.current?.value) return
        // check connection if not back to home page 
        if (conn===null){
            router.push('/')
            return
        }
        
        conn.send(textarea.current.value)
        textarea.current.value=''
    }
    
    return (
        <>
      <div className='flex flex-col w-full'>
        <div className='p-4 md:mx-6 mb-14'>
          <ChatBody data={messages} />
        </div>
        <div className='fixed bottom-0 mt-4 w-full'>
          <div className='flex md:flex-row px-4 py-2 bg-grey md:mx-4 rounded-md'>
            <div className='flex w-full mr-4 rounded-md border border-blue'>
              <textarea
                ref={textarea}
                placeholder='type your message here'
                className='w-full h-10 p-2 rounded-md focus:outline-none'
                style={{ resize: 'none' }}
              />
            </div>
            <div className='flex items-center'>
              <button
                className='p-2 rounded-md bg-blue text-white'
                onClick={sendMessage}
              >
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default index