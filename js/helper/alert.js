const alertMessage = (sender, Recipient ,div1) => {
    console.log("alertMessage");
    const div = document.createElement('div');
    div.className = 'notif';

    const span = document.createElement('span');
    span.className = 'welcome_text';

    span.innerHTML = `Hey, <span class="notif-user"> ${Recipient}</span>  you have a new message from <span class="notif-user"> ${sender}</span>`;

    div.appendChild(span);
    div1.appendChild(div);
}


export {alertMessage}