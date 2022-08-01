let table = document.createElement("table");
let header = table.createTHead();
let row = header.insertRow(0); 

let nameHeaderCell = document.createElement("th");
nameHeaderCell.innerHTML = "Name";
row.appendChild(nameHeaderCell);


let contactHeaderCell = document.createElement("th");
contactHeaderCell.innerHTML = "Contact";
row.appendChild(contactHeaderCell);

let i=1;
for(let key in localStorage) {
    if (!localStorage.hasOwnProperty(key)) {
      continue;
    }
    
    let row = table.insertRow(i)
    let nameCell = row.insertCell(0);
    let numberCell = row.insertCell(1);
    nameCell.innerHTML = key
    numberCell.innerHTML = localStorage.getItem(key)
    i++
}

let container = document.getElementById("content");
if(i>1) {
  container.appendChild(table)
}
else {
  let message = document.createElement("h2");
  message.innerHTML = "No contacts 😵"
  container.append(message)
  
}