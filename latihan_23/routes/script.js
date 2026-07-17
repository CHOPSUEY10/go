            

        


        class Event {

            constructor(type,payload){
                this.type = type
                this.payload = payload
            }


            routeEvent(){
                if (this.type === undefined){
                    alert("No type field in the event")
                }

                // FIX IT LATER FOR NEW FEATURES
            //     switch(this.type){
            //         case "new message" :
            //             console.log("new message")
            //             break
            //         default : 
            //             alert("unsupported message type")
            //             break     
            //     }
            }

            static sendEvent(eventName, payload){
                const event = new Event(eventName,payload)
                conn.send(JSON.stringify(event))
            }

        }

            const chatMessages = document.querySelector('#chatmessages');
            const MAX_LINES = parseInt(chatMessages.getAttribute('rows')) || 10;

            let selectedChat ="general"
            let conn = null
            const changeChatRoom  = (e) => {
                 e.preventDefault();
                let newchat = document.querySelector("#chatroom")
                if(newchat !== null && newchat.value != selectedChat){
                    console.log(newchat);
                }
                
                return false
                
            }   
            
            const sendMessage = (e) => {
                e.preventDefault();
                let newmessage = document.querySelector("#message")
            
                if(newmessage != null){
                    Event.sendEvent("send message",newmessage.value)
                }
                return false

            }

            const streamMessage =  (e) => {
            
                // A. Ambil isi teks yang sudah ada (termasuk isinya jika belum penuh)
                let newMessage =  e.payload
                let currentContent = chatMessages.value;


                // B. Tambahkan pesan baru di baris terbawah (APPEND)
                // JIKA textarea sudah ada isinya, tambahkan baris baru (\n) dulu
                if (currentContent !== "") {
                    chatMessages.value = currentContent + '\n' + newMessage;
                } else {
                    // JIKA masih kosong (misalnya baru connect), langsung isi
                    chatMessages.value = newMessage;
                }
                
                // --- Jaga Agar Scroll Di Paling Bawah ---
                chatMessages.scrollTop = chatMessages.scrollHeight;
            }

            window.onload = () => {
                document.querySelector("#chatroom-selection").onsubmit = changeChatRoom
                document.querySelector("#chatroom-message").onsubmit = sendMessage

                if(window["WebSocket"]){
                    
                    console.log("support websockets")
                    conn = new WebSocket("ws://" + location.host + "/ws");
                    
                    conn.onmessage = (e) => {
                        const eventData = JSON.parse(e.data)
                        const event = Object.assign(new Event,eventData)
                        event.routeEvent()    
                        streamMessage(eventData)
                    }

                    conn.onopen = () => console.log("CONNECTED");

                    conn.onerror = (e) => console.log("ERROR", e);

                    conn.onclose = (e) => console.log("CLOSED", e);
                }else{
                    alert("Browser does not support websocket")
                }
                
            }


