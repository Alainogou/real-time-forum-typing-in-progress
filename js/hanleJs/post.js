import { renderCommentForm } from "../componentsHtml/commentForm.js";
import { fetchComment, handleComment } from "./comment.js";


function fetchPost  (globalPosts,UserId) {
    globalPosts.innerHTML=''
    fetch('http://localhost:8081/fetchPost')
    .then(response => response.json())
    .then(response => {
        
        // postImage, friendName, postTime, postText, likeCount, commentCount, title, category

        for (let i=0; i<response.length;i++){

            let essai=document.createElement('div');
            const postHtml = fetchPosthtml(
                response[i].Post_id,
                './assets/imageUpload/'+response[i].ImageName,
                response[i].NickName,
                '16h.',
                response[i].Content,
                response[i].Nbrlike + ' Likes',
                response[i].NbrComments,      
                response[i].Title,
               
                response[i].Category,
                
            );
            essai.innerHTML=postHtml
            globalPosts.appendChild(essai)
            setupLikeButton(response[i].Post_id);

          
        }
      
        
        let commentButtons = document.querySelectorAll('.comment_btn');
       
        
        for (let i = 0; i < commentButtons.length; i++) {
           let commentButton = commentButtons[i];
           
            

           commentButton.addEventListener("click", (event) => {
               let postId = commentButton.querySelector('input[name="post_id"]').value;
               let addComment = document.querySelector(`.addComment_${postId}`);

               event.preventDefault();
               renderCommentForm(addComment, postId)
               
               if (addComment.style.display !== 'block') {
                   addComment.style.display = 'block';
               } else {
                   addComment.style.display = 'none';
               }

                let commentForms=document.querySelector(`.commentform-${postId}`)
                
               
                let containerComment=document.createElement('div')
                containerComment.classList.add("containerComment")
                commentForms.addEventListener('submit', function(event) {
                    
                    handleComment(event, UserId);
                    event.target.reset();
                    fetchComment(containerComment, postId)
                    
                  

                });
                

               
                fetchComment(containerComment , postId)
                addComment.appendChild(containerComment)
                
                
           });
        }
       
    
    })
    .catch(error => console.error('Erreur:', error));


}


function handleCreatePost(event, postform) {
    event.preventDefault();
    const formData = new FormData(event.target);
   
    let userid= parseInt(formData.get("user_id"))
    let postContent = {
        User_id: userid,    
        Title: formData.get("title"),
        Content: formData.get("content"),     
        Category: Array.from(formData.getAll("cat")).map(Number),
    }
    
    let file = document.querySelector('input[type="file"]').files[0];
    let reader = new FileReader();

    reader.onloadend = function() {
        let base64File = reader.result
        console.log(base64File)
        if (file) {
            postContent.Image = base64File;
        }

        fetch('http://localhost:8081/createPost', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json', 
            },
            body: JSON.stringify(postContent),
        })
        .then(response => {
            if (response.ok) {
                postform.style.display = 'none';
                let globalPosts=document.querySelector('.mainPost')
                fetchPost(globalPosts, userid)
                
            } else {
                return response.json();
            }
        })
        .then(errorResponse => {
            if (errorResponse) {
               
                
                switch (errorResponse ['error_class']) {
                    case 'categoryNofound':
                        showError(".messageErrorCategorie", errorResponse['message']);
                        break;
                    case 'titleNoFound':
                        showError(".messageErrorTitle", errorResponse['message']);
                        break;
                    case 'contentNofound':
                        showError(".messageErrorContent", errorResponse['message']);
                        break;
                    case 'imageNoCorrect':
                        showError(".messageErrorImage", errorResponse['message']);
                        break;
                    default:
                        console.error('Erreur inattendue:', errorResponse);
                }
            }
        })
        .catch(error => {
            console.error('Erreur lors de la création de l\'utilisateur:', error);
        });
    };

    if (file) {
        reader.readAsDataURL(file);
    } else {
        reader.onloadend();
    }
}

function showError(selector, message) {
    let errorElement = document.querySelector(selector);
    errorElement.innerHTML = message;
    setTimeout(() => {
        errorElement.innerHTML = '';
    }, 5000);
}


function fetchPosthtml(postId, postImage, friendName, postTime, postText, likeCount, commentCount, title, category) {
   
    let categoryHtml = category.map(cat => `<h2>${cat}</h2>`).join('');
    let imageHtml = (postImage!=='./assets/imageUpload/') ? `<img src="${postImage}">` : '';

    return `
      <div class="friends_post">
          <div class="friend_post_top">
              <div class="img_and_name">
                 
              <div class="user" style="background-color:#efefef; height:30px;width:30px; text-align:center; border-radius:50%; padding-top:4px">
              <i class="fa-solid fa-user" ></i>
              </div>
             
                  <div class="friends_name">
                      <p class="friends_name">
                          ${friendName}
                      </p>
                      
                      <p class="time">${postTime}<i class="fa-solid fa-user-group"></i></p>
                  </div>
              </div>
              <div class="menu">
                   
                    ${categoryHtml}
            
                 
              </div>

          </div>
          
          <div>
             <h1 style="font-size:20px">${title}</h1>
             <br></br>
              <p class="postText">${postText}</p>
              <br></br>
          </div>
          ${imageHtml}
          <div class="info">
              <div class="emoji_img">
                  <img src="/assets/image/like.png">
                  <img src="/assets/image/haha.png">
                  <img src="/assets/image/heart.png">
                  <p id="likeCount-${postId}">${likeCount} </p>
              </div>
              <div class="comment">
                  <p ><span id="commentNumber-${postId}">${commentCount} </span> Comments</p>
                
              </div>
          </div>
          <hr>
          <div class="like">
          <div class="like_icon" id="likeButton-${postId}" data-liked="false">
          <i class="fa-solid fa-thumbs-up activi"></i>
                  <p>Like</p>
              </div>
              
           
              <div class="like_icon comment_btn">
                    <input type="hidden" name="post_id" value="${postId}">
                    <i class="fa-solid fa-message"></i>
                    <p>Comments</p>
             </div>
            
          </div>
          
          <div class="addComment_${postId}">

          </div>
         

      </div>
      
      `;
  }



// Cette fonction configure le bouton "J'aime" pour un post spécifique.
function setupLikeButton(postId) {
    var likeButton = document.getElementById('likeButton-' + postId);
    var likeCountElement = document.getElementById('likeCount-' + postId);

    // Récupérer l'état "aimé" et le nombre de likes du localStorage
    var isLiked = localStorage.getItem('liked-' + postId) === 'true';
    var likeCount = parseInt(localStorage.getItem('likeCount-' + postId)) || 0;

    // Mettre à jour l'interface utilisateur avec les valeurs récupérées
    likeCountElement.textContent = likeCount + ' Likes';
    likeButton.setAttribute('data-liked', isLiked.toString());
    likeButton.classList.toggle('liked', isLiked);

    // Ajouter un écouteur d'événements pour gérer les clics sur le bouton "J'aime"
    likeButton.addEventListener('click', function() {
        isLiked = !isLiked;
        likeCount = isLiked ? likeCount + 1 : likeCount - 1;

        // Mettre à jour l'interface utilisateur
        likeCountElement.textContent = likeCount + ' Likes';
        likeButton.setAttribute('data-liked', isLiked.toString());
        likeButton.classList.toggle('liked', isLiked);

        // Mettre à jour le localStorage avec le nouvel état et le nouveau nombre de likes
        localStorage.setItem('liked-' + postId, isLiked.toString());
        localStorage.setItem('likeCount-' + postId, likeCount.toString());
    });
}

// Appeler setupLikeButton pour chaque post lorsque la page est chargée.
document.addEventListener('DOMContentLoaded', function() {
    var likeButtons = document.querySelectorAll('[id^="likeButton-"]');
    likeButtons.forEach(function(button) {
        var postId = button.id.split('-')[1];
        setupLikeButton(postId);
    });
});

  export {fetchPost, handleCreatePost}