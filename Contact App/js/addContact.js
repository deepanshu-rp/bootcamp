const KEY = "contacts"

function addContact(form) {
    let contact = {
        "name": form.elements["name"].value,
        "number": form.elements["contact_number"].value
    }

    let oldContacts = localStorage.getItem(KEY)
    if(oldContacts === null) {
        let newContacts = [contact]
        localStorage.setItem(KEY, JSON.stringify(newContacts))
    } else {
        let newContacts = JSON.parse(oldContacts)
        newContacts.push(contact)
        localStorage.setItem(KEY, JSON.stringify(newContacts))
    }
    alert("Contact Saved!")
}