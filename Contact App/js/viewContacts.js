let container = document.getElementById("content");
let storedContacts = localStorage.getItem("contacts")

if(storedContacts === null) {
  let htmlMessageString = "<h2>No contacts 😵</h2>"
  container.insertAdjacentHTML("afterbegin", htmlMessageString)
}
else {
  let contacts = JSON.parse(storedContacts)
  let htmlRowsString = "<tr><th>Name</th><th>Number</th></tr>"
  
  for(let i in contacts) {
    htmlRowsString += `<tr><td>${contacts[i].name}</td><td>${contacts[i].number}</td></tr>`
  }
  
  let htmlTableString = "<table>" + htmlRowsString + "</table>";
  container.insertAdjacentHTML("afterbegin", htmlTableString)
}