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
		case cl:=<-h.Register:
			//the room exist
			if _ , ok := h.Rooms[cl.RoomID]; ok{
			  	 r:=h.Rooms[cl.RoomID]

				 //the client if  is already  exist or not 
				 //if doesn't exist then add client to the room
				 if _,ok := r.Clients[cl.ID]; !ok{
					r.Clients[r.ID] = cl
				 }
			}
		case cl:=<-h.UnRegister:
			//check the roomId exist
			if _,ok :=  h.Rooms[cl.RoomID] ; ok{
				//if the client actually exists
				if _,ok:= h.Rooms[cl.RoomID].Clients[cl.ID] ; ok{
					//broadcast that the client has left the room
					if len(h.Rooms[cl.RoomID].Clients)!=0{
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}
					//client leave the room
					delete(h.Rooms[cl.RoomID].Clients , cl.ID)
					close(cl.Message)
				}
			} 
		case m := <-h.Broadcast: 
			//check the the exist or not
			if _ , ok :=  h.Rooms[m.RoomID] ; ok{
				//sending the message to all the client in the room
				for _ , cl := range h.Rooms[m.RoomID].Clients{
					cl.Message<- m 
				}
			}
		}
	}
}