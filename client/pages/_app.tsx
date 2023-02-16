import '../styles/globals.css'
import type { AppProps } from 'next/app'
import AuthContextProvider from '../module/auth_provider'
import WebSocketProvider from '../module/websocket_provider'

export default function App({ Component, pageProps }: AppProps) {
  return (
  <>
    <AuthContextProvider>
      <WebSocketProvider>
      <div className='flex flex-col md:flex-row h-full min-h-screen font-sans'>
        <Component {...pageProps}></Component>
      </div>
      </WebSocketProvider>
    </AuthContextProvider>
  </>
  )
}
