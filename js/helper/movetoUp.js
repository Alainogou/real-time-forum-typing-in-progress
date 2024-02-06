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

export {moveUserToTop}