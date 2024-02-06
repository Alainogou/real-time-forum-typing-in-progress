function handleComment(event, userId){
    event.preventDefault();
    
    const formData = new FormData(event.target);
    let postId=parseInt(formData.get("post_id"))
    let newComment={
        UserId :parseInt(userId),
        Content: formData.get("content"),
        Post_id:postId,
        
    } 
    
    fetch('http://localhost:8081/createComment', {
        method: 'POST',
        headers: {
           'Content-Type': 'application/json', 
       },
       body: JSON.stringify(newComment),
     
    })
   .then(response => {
       if (response.ok) {
            console.log(newComment);
            let commentNumber= document.getElementById(`commentNumber-${postId}`)
            commentNumber.innerText= parseInt(commentNumber.textContent ) + 1 
           
        } else {       

            return response.json();
        }
    })
    .then(response => { 
       
       if (response){
       
        let emptyContent= document.querySelector(".EmptyContent")
        if (response['error_class']==="emptycomment"){
            emptyContent.innerHTML=response['message']
        }
        setTimeout(() => {
            emptyContent.innerHTML=''
        }, 5000);
    }
   
       

    })
 
   .catch(error => console.error('Erreur lors de la création de l\'utilisateur:', error));
  
}


function fetchComment(addcomment, postId){
  

    fetch(`http://localhost:8081/fetchComment/${postId}`)
    .then(response => response.json())
    .then(data => {

        
        addcomment.innerHTML=''
        

        for (let p=0;p<data.length;p++){
            let comment=data[p]

            let mainComment= document.createElement('div')
            mainComment.classList.add("mainComment")


            let commentProfile= document.createElement('div')
            commentProfile.classList.add("commentProfile")

            let iconeProfile= document.createElement('div')
            iconeProfile.innerHTML=`<div class="user" style="background-color:#efefef; height:30px;width:30px; text-align:center; border-radius:50%; padding-top:4px">
                                        <i class="fa-solid fa-user" ></i>
                                     </div>`
            iconeProfile.classList.add("iconeProfile")
        
            let userComment= document.createElement('p')
            userComment.classList.add("userComment")
            userComment.innerText=`${comment.NickName} :`
            commentProfile.append(iconeProfile,userComment);
            // commentProfile.appendChild(userComment);

            let contentComment= document.createElement('div')
            let pContent= document.createElement('p')

            contentComment.classList.add("contentComment")

            pContent.innerText=`${comment.Content}`
            contentComment.appendChild(pContent)
            mainComment.appendChild(commentProfile)
            mainComment.appendChild(contentComment)


            
            addcomment.appendChild(mainComment)
            
        }
 
        
    
    })
    .catch(error => console.error('Erreur:', error));


}

export {handleComment, fetchComment}