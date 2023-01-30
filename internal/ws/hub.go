package ws

type Room struct{
	ID		string  `json:"id"`
	Name 	string	`json:"name"`
	Clients	map[string]*Client	`json:"clients"`
}

type Hub struct{
   	 Rooms 	 	map[string]*Room
	 Register	chan *Client
	 UnRegister	chan *Client
	 Broadcast	chan *Message
}

func NewHub() *Hub{
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Broadcast:  make(chan *Message , 5),
	}
}

func (h *Hub) Run() {
	for {
		select{
		case client :=<- h.Register:
			//the room exist
			if _ , isRoomExist := h.Rooms[client.RoomID]; isRoomExist{

				 //the client if  is already  exist or not 
				 //if doesn't exist then add client to the room
				 room := h.Rooms[client.RoomID]
				 
				 if _,isClientExist := room.Clients[client.ID]; !isClientExist{
					room.Clients[client.ID] = client
				 }
			}
			
		case client := <-h.UnRegister:
			//check the roomId exist
			if _,ok :=  h.Rooms[client.RoomID] ; ok{
				//if the client actually exists
				if _,ok:= h.Rooms[client.RoomID].Clients[client.ID] ; ok{
					//broadcast that the client has left the room
					if len(h.Rooms[client.RoomID].Clients)!=0{
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   client.RoomID,
							Username: client.Username,
						}
					}
					//client leave the room
					delete(h.Rooms[client.RoomID].Clients , client.ID)
					close(client.Message)
				}
				//if the room dosen't exist then delete the room
				clients := h.Rooms[client.RoomID].Clients
				if len(clients) == 0 {
					delete(h.Rooms, client.RoomID)
				}
			} 
		case message := <-h.Broadcast: 
			//check the the exist or not
			if _ , ok :=  h.Rooms[message.RoomID] ; ok{
				//sending the message to all the client in the room
				for _ , client := range h.Rooms[message.RoomID].Clients{
					
					client.Message<- message
					
				}
			}
		}
	}
}