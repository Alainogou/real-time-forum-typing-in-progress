


function chatContainer(toUser, fromUser, premierDiv)  {
    
    const socket = new WebSocket(`ws://localhost:8081/communication?userFrom=${fromUser}&toUser=${toUser}`);
   
       // Écoutez les messages entrants
    socket.onmessage = function(event) {
        
        let msg = JSON.parse(event.data);
        
        FormMessage(premierDiv, toUser,fromUser, msg.UserReceiver, msg.UserForum);
        

        sendMessage(toUser, fromUser, premierDiv)
        
    };
    
             
    socket.onopen = (event) => {
       
    };

   

    socket.onclose = () => {
        console.log('WebSocket connection closed');
    };

    socket.onerror = (error) => {
        console.log(`WebSocket error: ${error}`);
    };

}


function sendMessage(toUser, fromUser, premierDiv)  {
    

    const socket = new WebSocket(`ws://localhost:8081/privateSocket?userFrom=${fromUser}&toUser=${toUser}`);

    socket.onopen = (event) => {

        let closeMesenger= document.querySelector(`.btn-close2-${toUser}`)
        closeMesenger.addEventListener("click",()=>{
                let clickClose = document.querySelector(`.chat-card-${toUser}`)
                clickClose.remove()
                const messageData = {
                    
                    ToUserClosed:fromUser+toUser
                };
                socket.send(JSON.stringify(messageData))

        })
      
        let typingTimeout;
        const messageInput = document.querySelector(`input[name="messagePrivite-${toUser}"]`)
        messageInput.addEventListener('keydown', () => {
           
            if (typingTimeout) clearTimeout(typingTimeout);
        
            socket.send(JSON.stringify({IsTyping: true }));
            
            typingTimeout = setTimeout(() => {
                socket.send(JSON.stringify({IsTyping: false }));
            },  1000); 
        });

        messageInput.addEventListener('keyup', () => {
           
            if (typingTimeout) clearTimeout(typingTimeout);
          
            typingTimeout = setTimeout(() => {
                socket.send(JSON.stringify({IsTyping: false }));
            },  1000); 
        });

        messageInput.addEventListener('blur', () => {
            socket.send(JSON.stringify({IsTyping: false }));
        });


        let messageFormId = document.querySelector(`.receved-${toUser}`)
       
        if (messageFormId)  messageFormId.addEventListener('submit', (event) => {
            event.preventDefault();
            const messageInput = document.querySelector(`input[name="messagePrivite-${toUser}"]`);

            const message = messageInput.value;
            let emptyMessage = document.querySelector(".emptyMsg")
            if ( message=="") {
                emptyMessage.textContent="Message empty"
                emptyMessage.style.color="red"
                setTimeout(() => {
                    emptyMessage.textContent=""
                }, 3000);
                console.log("Message empty");
            }
            const messageData = {
                FromUser: fromUser,
                ToUser: toUser,
                Message: message,
                CreateDate: new Date().toISOString()
            };
            
            socket.send(JSON.stringify(messageData));
            messageInput.value = ''; // Effacer le champ de saisie après l'envoi
            
            if(message!==""){
                moveUserToTop(toUser)
            }
            
            
          
        });


    };
    
    socket.onmessage = function(event) {
        
        let msg = JSON.parse(event.data);
       
        console.log(msg, "MON NESSAE")
        
        FormMessage(premierDiv, toUser,fromUser, msg.UserReceiver, msg.UserForum);

         
        let typingTimeout;
        const messageInput = document.querySelector(`input[name="messagePrivite-${toUser}"]`)
        messageInput.addEventListener('keydown', () => {
           
            if (typingTimeout) clearTimeout(typingTimeout);
        
            socket.send(JSON.stringify({IsTyping: true }));
            
            typingTimeout = setTimeout(() => {
                socket.send(JSON.stringify({IsTyping: false }));
            },  1000); 
        });

        messageInput.addEventListener('keyup', () => {
           
            if (typingTimeout) clearTimeout(typingTimeout);
          
            typingTimeout = setTimeout(() => {
                socket.send(JSON.stringify({IsTyping: false }));
            },  1000); 
        });

        messageInput.addEventListener('blur', () => {
            socket.send(JSON.stringify({IsTyping: false }));
        });


        let closeMesenger= document.querySelector(`.btn-close2-${toUser}`)

        const typingNotification = document.querySelector('.typing-notification');
        if (msg.IsTyping){
             typingNotification.style.display = 'block';
        }else{
            typingNotification.style.display = 'none';
        }
       
        closeMesenger.addEventListener("click",()=>{
                let clickClose = document.querySelector(`.chat-card-${toUser}`)
                clickClose.remove()
                const messageData = {
                    
                    ToUserClosed:fromUser+toUser
                };
                socket.send(JSON.stringify(messageData))
        })
        
    
        
        let messageFormId = document.querySelector(`.receved-${toUser}`)
        if (messageFormId)  messageFormId.addEventListener('submit', (event) => {
            event.preventDefault();
            const messageInput = document.querySelector(`input[name="messagePrivite-${toUser}"]`);
            const message = messageInput.value;
            const messageData = {
                FromUser: fromUser,
                ToUser: toUser,
                Message: message,
                CreateDate: new Date().toISOString()
            };
            
            socket.send(JSON.stringify(messageData));
            messageInput.value = ''; // Effacer le champ de saisie après l'envoi
            if(message!==""){
                moveUserToTop(toUser)
            }
          
        });

        
       
    };
    

    socket.onclose = () => {
        console.log('WebSocket connection closed');
    };

    socket.onerror = (error) => {
        console.log(`WebSocket error: ${error}`);
    };

}


const FormMessage = (premierDiv, toUser, userFrom, messageRecu, messageEnvoyer) => {
    // Initialize the chat body
    premierDiv.innerHTML=""
    let spinner=document.createElement('div');
    spinner.classList.add('spinner');
    spinner.classList.add('center')
   
    spinner.style.display='none'

    let chatBody = document.createElement('div');
    chatBody.className = 'chat-body';
     

    let allMessages = messageRecu.concat(messageEnvoyer);

    allMessages.sort((a, b) => new Date(a.CreateDate) - new Date(b.CreateDate));
    
    if (allMessages.length<=10){
       
        loadMessages(0,allMessages.length, chatBody, allMessages, userFrom);
    }else{
        
        let messages=[]
        messages.push(...allMessages);
       
        var tampon= [];
        var countTamp=0;
       
        // Get the last 10 message
        for (let t = 0; t <  messages.length; t++) {
           
            if (messages[t]) {                            
                    let tamp = allMessages.pop();
                    tampon.push(tamp);
            }
           
            countTamp++;
            if (countTamp==10) break
        }
        var tamponSecond= [];
        for (let p=tampon.length-1;p>=0;p--){
            let tamp = tampon.pop();
            tamponSecond.push(tamp);
        }
        messages=[]
        messages.push(...allMessages);
       
      
        loadMessages(0, 10, chatBody, tamponSecond, userFrom)
        chatBody.scrollTop = chatBody.scrollHeight;
        
        var countLoader=0;
        chatBody.addEventListener('scroll',(event)=>{
            if (event.target.scrollTop === 0) {
                let scrollHeightBefore = chatBody.scrollHeight;
                if (countLoader==0) {   
                    spinner.style.display='block'                     
                    countLoader++;
                }

                setTimeout(() => {                        
                    spinner.style.display='none'
                    countLoader=0;
                }, 499);
            
                let scrollHeightAfter = chatBody.scrollHeight;
                chatBody.scrollTop = chatBody.scrollTop + (scrollHeightAfter - scrollHeightBefore);
            }
        });

        chatBody.addEventListener('scroll', throttle((event) => {
            // Check if user has scrolled to the top
            if (event.target.scrollTop === 0) {
                let scrollHeightBefore = chatBody.scrollHeight;
        
                for (let t = 0; t < Math.min(10,messages.length-1); t++) {
                    let _data = allMessages.pop();

                    if (_data) {                          
                                            
                            let newMessage;
                            newMessage=AddMessage(_data.FromUser, _data.ContentMessage, userFrom )
                            let newTime= AddTime(timeAgo(_data.CreateDate))
                            chatBody.insertBefore(newTime, chatBody.firstChild);
                            chatBody.insertBefore(newMessage, chatBody.firstChild);
                           
                        }
                }
                // Adjust scroll position to prevent jumping
                messages=[]
                messages.push(...allMessages);
                console.log("apres scroll", messages.length)
                let scrollHeightAfter = chatBody.scrollHeight;
                chatBody.scrollTop = chatBody.scrollTop + (scrollHeightAfter - scrollHeightBefore);
            }
        }, 500));


    }

     // Create the chat header
     let chatHeader = document.createElement('div');
     chatHeader.className = 'chat-header';
     chatHeader.innerHTML = `
         <div class="h2">
            <div class="profChat">
             <div class="user" style="background-color:#efefef; height:30px;width:30px; text-align:center; border-radius:50%; padding-top:4px">
                 <i class="fa-solid fa-user"></i>   
             </div>
             <p>${toUser}</p>
             </div>
             <div class="typing-notification noty">
                 <span>${toUser} is writing...</span> 

                 <div class="newtons-cradle">
                 <div class="newtons-cradle__dot"></div>
                 <div class="newtons-cradle__dot"></div>
                 <div class="newtons-cradle__dot"></div>
                 <div class="newtons-cradle__dot"></div>
                 </div>
                
             </div>
             <button class="btn-close2 btn-close2-${toUser}">X</button>
         </div>
     `;

    
 
    //  let notif=document.createElement('div')
    //  notif.className = 'typing-notification';
    //  notif.textContent=`${toUser} is writing...`
     
     // Create the chat foot
     let chatFooter = document.createElement("div");
     chatFooter.innerHTML=` 
             <form enctype="multipart/form-data" class="receved-${toUser}">
                <input type="hidden" name="send-Name" value="${userFrom}">
               
                <div class="chat-footer"> 
                <p class="emptyMsg"></p>

                    <input placeholder="Type your message" type="text" name="messagePrivite-${toUser}">
                    <button>Send</button>
                </div>
            </form>
            
    `
 
     
     // Append everything to the premierDiv
    
     premierDiv.appendChild(chatHeader);

     premierDiv.appendChild(spinner)
     premierDiv.appendChild(chatBody);
    //  premierDiv.appendChild(notif)
     premierDiv.appendChild(chatFooter);
    //  right.appendChild(premierDiv);
    
 
};

function throttle(func, limit) {
    let lastFunc;
    let lastRan;
    return function() {
        const context = this;
        const args = arguments;
        if (!lastRan) {
            func.apply(context, args);
            lastRan = Date.now();
        } else {
            clearTimeout(lastFunc);
            lastFunc = setTimeout(function() {
                if ((Date.now() - lastRan) >= limit) {
                    func.apply(context, args);
                    lastRan = Date.now();
                }
            }, limit - (Date.now() - lastRan));
        }
    }
}

function loadMessages (startIndex , limit, chatBody , messages, userFrom) {

       
    for (let i = startIndex; i <  limit; i++) {
        let message = messages[i];
        if (message){
            let messageElement = document.createElement('div');
            messageElement.classList.add( message.FromUser === userFrom ? 'outgoing' : 'incoming');
    
            messageElement.classList.add('messageSMS')
            let contentMessage = document.createElement('div');
            contentMessage.classList.add("contentMessage")
    
            let contentElement = document.createElement('p');
            contentElement.classList.add( "message-conten");
            contentElement.innerText = `${message.ContentMessage}`;
            let userElement = document.createElement('span');
            userElement.classList.add("message-sender")
            userElement.innerText = `${message.FromUser}`;
    
            let timeElement = document.createElement('div');
            timeElement.classList.add("message-date")
            let time=timeAgo(message.CreateDate)
            timeElement.innerText = `${time}`;
    
        
            contentMessage.appendChild(contentElement)
            contentMessage.appendChild(userElement)
            messageElement.appendChild(contentMessage);
            
          
            chatBody.appendChild(messageElement);
            chatBody.appendChild(timeElement);
        }
       
     
    }
};

function AddMessage(sender, content, userFrom){
            let messageElement = document.createElement('div');
            messageElement.classList.add( sender === userFrom ? 'outgoing' : 'incoming');
    
            messageElement.classList.add('messageSMS')
            let contentMessage = document.createElement('div');
            contentMessage.classList.add("contentMessage")
    
            let contentElement = document.createElement('p');
            contentElement.classList.add( "message-conten");
            contentElement.innerText = `${content}`;
            let userElement = document.createElement('span');
            userElement.classList.add("message-sender")
            userElement.innerText = `${sender}`;
    
            contentMessage.appendChild(contentElement)
            contentMessage.appendChild(userElement)
            messageElement.appendChild(contentMessage);
            
            return messageElement
}

function AddTime(time){

    let timeElement = document.createElement('div');
    timeElement.classList.add("message-date")
    timeElement.innerText = `${time}`
    
    return timeElement
}

function timeAgo(dateString) {
    const date = new Date(dateString);
    const now = new Date();
  
    const secondsPast = (now.getTime() - date.getTime()) / 1000;
  
    if(secondsPast < 60) {
        return parseInt(secondsPast) + ' sec ago';
    }
  
    if(secondsPast < 3600) {
        return parseInt(secondsPast/60) + ' mn ago';
    }
  
    if(secondsPast <= 86400) {
        return parseInt(secondsPast/3600) + ' h ago';
    }
  
    if(secondsPast > 86400) {
        const daysPast = parseInt(secondsPast/86400);
        if (daysPast < 7) {
            return daysPast + ' days ago';
        } else if (daysPast < 30) {
            return parseInt(daysPast/7) + ' weeks ago';
        } else if (daysPast < 365) {
            return parseInt(daysPast/30) + ' months ago';
        } else {
            return parseInt(daysPast/365) + ' years ago';
        }
    }
}


function moveUserToTop(userId) {
    console.log("top");
    const usersDiv = document.querySelector(".userlist");
    const userContainer = document.querySelector(`.contact-${userId}`);
    if (userContainer) {
      console.log("userContainer: ", userContainer);
      usersDiv.removeChild(userContainer);
      usersDiv.insertBefore(userContainer, usersDiv.children[0]);
      // Enregistrer l'ordre des utilisateurs dans localStorage
    
    }
}
export {chatContainer}