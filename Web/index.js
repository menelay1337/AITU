'use strict'
// modal script
//
var modal = document.getElementById("myModal");

var btn = document.getElementById("modalBtn");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];
btn.onclick = function() {
  modal.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
  modal.style.display = "none";
}


// image slider

var slides = document.querySelectorAll('.slide');
var currentSlide = 0;
var slideInterval = window.setInterval(nextSlide,20000);

function nextSlide() {
  slides[currentSlide].className = 'slide';
  currentSlide = (currentSlide+1)%slides.length;
  slides[currentSlide].className = 'slide active';
}

function previousSlide() {
  slides[currentSlide].className = 'slide';
  currentSlide = (((currentSlide-1)%slides.length)+slides.length)%slides.length;
  slides[currentSlide].className = 'slide active';
}

nextSlide();

// login script
//

//function changeForm(){
//	if (form.formbutton.value === 'Have an account.') {
//		form.formsumbit.onclick = loginValidate;
//		form.formsumbit.value = "Login";
//		form.formbutton.value = "Need to register.";
//		document.getElementById("logheader").innerHTML = "Login:"
//		let header = document.getElementById("formHeader");
//		header.innerHTML = "Login";
//		document.getElementById("userName").innerHTML = "Login";
//	} else if ( form.formbutton.value === "Need to register.") {
//		form.formsumbit.onclick = registerValidate;
//		form.formsumbit.value = "Register";
//		form.formbutton.value = "Have an account.";
//		document.getElementById("logheader").innerHTML = "New login:"
//		let header = document.getElementById("formHeader");
//		header.innerHTML = "Register";
//		document.getElementById("userName").innerHTML = "Register";
//	} else {
//		window.alert("Logic error!");
//	}
//}


document.getElementById("myForm").addEventListener('submit', function(event) {
	event.preventDefault();
	registerValidate();
});


function quit(){
		document.getElementById("userName").innerHTML = "Register";
		document.querySelector(".panel").style.display = "none";
		window.alert("You left from account")
		btn.onclick = function() {
		  modal.style.display = "block";
		}
		
      		
}

var users = []
function registerValidate(){
	var login = document.forms.myForm.login.value;
	let email = document.forms.myForm.email.value;
	let passwd = document.forms.myForm.passwd.value;
	if (login.length > 10) {
		window.alert("login length too long");
		
	} else if (passwd.length < 8 ){
		window.alert("password too weak!")
	} else if ( sessionStorage.getItem(login) === null ) {
		sessionStorage.setItem(login, passwd);
		users.push({"login": login, "email": email, "passwd": passwd});
		window.alert("Account has registered!");
		modal.style.display = "none";
		if (login === "Admin") {
			let json = JSON.stringify(users);
			sessionStorage.setItem("users", json);
			document.querySelector(".panel").style.display = "block";
			window.alert("Admin panel has opened!");
		}
		document.getElementById("userName").innerHTML = login;
		btn.onclick = quit;
	} else if ( passwd === sessionStorage.getItem(login) ) {
		window.alert("Successful authorizing");
		modal.style.display = "none";
		if (login === "Admin") {
			let json = JSON.stringify(users);
			sessionStorage.setItem("users", json);
			document.querySelector(".panel").style.display = "block";
			window.alert("Admin panel has opened!");
		}
		document.getElementById("userName").innerHTML = login;
		btn.onclick = quit;
	} else {
		window.alert("Authorization error!");
	}
}

//function loginValidate(){
//	var login = document.forms.myForm.login.value;
//	let passwd = document.forms.myForm.passwd.value;
//	let currentDoc = document;
//
//	if (login.length > 10) {
//		window.alert("login length too long");
//	} else if ( sessionStorage.getItem(login) != null ) {
//		let storagePass = sessionStorage.getItem(login);
//		if (  storagePass === passwd ){
//			window.alert("Logged succesfully!");
//			modal.style.display = "none";
//			if (login === "Admin") {
//				document.querySelector(".panel").style.display = "block";
//				window.alert("Admin panel has opened!");
//			}
//			document.getElementById("userName").innerHTML = login; 
//			btn.onclick = quit;
//			window.document = currentDoc;
//
//		} else {
//			window.alert("Error while authorization");
//		}
//		//window.history.go (-1);
//	} else {
//		window.alert("This account doesn't exist!");
//	}
//}


function trackIcon(){
	icon.onmouseover = mouseOver;
	icon.onmouseout = mouseOut;
}

function mouseOver() {
  myImage.src = "my_image_active.png";
}
function mouseOut() {
  myImage.src = "my_image.png";
}


// counter script
//
var thumbsCount = 0;

function startCounting() {
	window.setInterval(plusOne, 1000);
	let element = document.getElementById("btnStart");
	element.remove();
}


function plusOne() {
	++thumbsCount;
	document.getElementById("thumbsCounter").innerHTML = thumbsCount;
}

