function validateInfo(contact_name, contact_number) {
    let phone = contact_number;
    var phoneNum = /^\(?(\d{3})\)?[- ]?(\d{3})[- ]?(\d{4})$/; 
    return contact_name.length>0 && phone.match(phoneNum)
}


function addContact(form) {
    let contact_name = form.elements["name"].value
    let contact_number = form.elements["contact_number"].value

    if(!validateInfo(contact_name, contact_number)) {
        alert("Invalid input")
        return
    }

    localStorage.setItem(contact_name, contact_number)
    alert("Contact Saved!")
}