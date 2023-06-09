function openHeader(){
    if(document.querySelector('#home').classList.contains('open')){
        let time = document.querySelector('#header.open').offsetWidth;
        time = time / 700 * 2000;
        setTimeout(() => {
            document.querySelector('#home').classList.toggle('open')
            document.querySelector('#circle').classList.toggle('open')
        }, time);  
    }
    else{
        document.querySelector('#home').classList.toggle('open')
        document.querySelector('#circle').classList.toggle('open')
    }
    document.querySelector('#header').classList.toggle('open')
}


let hiddenButton = document.querySelector(".hiddenButton")
let buttonForPhotos = document.querySelector(".buttonForPhotos")
let nothing = document.querySelectorAll('.inputPhotos')[0].files;  // lastChange   */
let inputPhotos = document.querySelector('.inputPhotos');  

let PHOTOS = {
    photos: [],
    start:0,
};


function showHiddenButton(PHOTOS){        
    
    for(let i = PHOTOS.start;i < PHOTOS.photos.length;i++){
        let elem = document.querySelector('.Photos')
        elem.append(tmpl.content.cloneNode(true));
    }
    PHOTOS.start = PHOTOS.photos.length;
    
    document.querySelectorAll('.photo').forEach((photo, i) => {
        photo.style.backgroundImage = `url(${URL.createObjectURL(PHOTOS.photos[i])})`
        // URL.revokeObjectURL(PHOTOS.photos[i]);  
    })
    
    document.querySelectorAll('.inputPhotos')[0].files = nothing; // сброс последнего изменения => если был удалена фотка 
                                                    // и пользователь захочет снова добавить эту фотку, он сможет это сделать, т.к. onchange сработает
    

    buttonForPhotos.classList.remove('open')
    hiddenButton.classList.add('open') 


}
function removeHiddenButton(){
    buttonForPhotos.classList.add('open')
    hiddenButton.classList.remove('open') 
}


function onclickDelete(event){
    document.querySelectorAll('.delete').forEach((item, i) =>{
        if(item === event){
            PHOTOS.photos.splice(i, 1);
            document.querySelectorAll('.photo')[i].remove();
            
            PHOTOS.start--;
            if(PHOTOS.photos.length == 0) removeHiddenButton();
        }
    })  

}

inputPhotos.oninput  = function(){
   let photos = document.querySelectorAll('.inputPhotos')[0];
   for(let i = 0; i < photos.files.length;i++){
       if(PHOTOS.photos.length >= 10){
           alert("максимальное кол-во файлов")
           break;
       }    
       PHOTOS.photos.push(photos.files[i]);
   }
   showHiddenButton(PHOTOS);
}



function deleteChild() {
document.querySelectorAll(".emotion").forEach((elem, i) => {
    elem.remove();
})
document.querySelectorAll('.recommendations').forEach(elem => {
    elem.remove();
})
document.querySelectorAll(".noemotion").forEach((elem, i) => {
    elem.remove();
})
}


function addResponse(response){
for(let i = 0; i < response.length; i++){
    if(response[i].emotion){
        let elem = document.querySelector('#block3')
        elem.append(tmp2.content.cloneNode(true));      
    }
    else{
        let elem = document.querySelector('#block3')
        elem.append(tmp3.content.cloneNode(true));
    }
}
}

function addEmotion(emotion, textEmotion){
let t = document.createElement('font');  
t.style.color = '#00E5E5';
t.style.fontWeight = "600"
t.innerHTML = ` ${textEmotion}`;
emotion.append(t);
}

function addRecommendationIMG(recommendations, photos){
let img = document.createElement('img');  
img.style.width = "220px"
img.style.float = "right"
img.style.marginTop = "-180px"
img.style.padding = "20px"
if(photos){
    img.src = `${URL.createObjectURL(photos)}`;
    recommendations.append(img)
}
}
function addTextRecommendation(recommendations, text){
let r = document.createElement('p');  
r.style.fontWeight = "500"
r.innerHTML = text;
recommendations.append(r);
}

function  addNoEmotion(recommendations, photos){
let img = document.createElement('img');  
img.style.width = "200px"
img.style.marginTop = "40px"
img.style.padding = "20px"
img.style.float = "right"
if(photos){
    img.src = `${URL.createObjectURL(photos)}`;
    recommendations.append(img)
}
}





let idLoading;
let numbPoint = 9;  
function loading(){
if(idLoading){
    deleteLoading();
}
let elem = document.querySelector('.expectation')
for(let i = 0; i < numbPoint; i++){
    elem.append(tmp4.content.cloneNode(true))
}
let radius = 8;
let speed = 20;
let f = [];
let r = 0;
for(let i = 0; i < numbPoint; i++){
    f[i] = r * Math.PI / 180;
    r += 360 / numbPoint;
}

let s = 2 * Math.PI / 180;

let tmpPoints = document.querySelectorAll('.point');
tmpPoints.forEach((point, i)=> {
    point.style.opacity = (numbPoint / (i + 1) / 2).toFixed(4)
})
idLoading = setInterval(() => {
    for(let i = 0; i < numbPoint; i++){
        tmpPoints[i].style.marginLeft = `${230 + radius * Math.sin(f[i])}px`
        tmpPoints[i].style.marginTop = `${-12 + radius * Math.cos(f[i])}px`
        tmpPoints[i].style.transform = `rotate(${-(f[i]) * 180 / Math.PI}deg)`
        f[i]+=s;
    }
}, 30);
setTimeout(() => elem.style.display = "block", 30);

}
function deleteLoading(){
document.querySelector('.expectation').style.display = "none"
clearInterval(idLoading);
document.querySelectorAll('.point').forEach((point, i) => {
    point.remove();
})
}





function sendPhotos() {

    loading();


deleteChild();

let formData = new FormData();
for(let i = 0; i < PHOTOS.photos.length; i++){
    console.log(PHOTOS.photos[i]);
    formData.append("file_names", PHOTOS.photos[i].name);
    formData.append("data", PHOTOS.photos[i]);
}

fetch('http://localhost:8080', {
    method: 'POST',
    body: formData,
})
.then(response => {
    if(response.ok){
        return response.json();
    }
})
.then(res => {
    res = res.PhotoDescription
    let photos = PHOTOS.photos;
    addResponse(res)

    let emotions =  document.querySelectorAll('.emotion');
    let recommendations = document.querySelectorAll('.recommendations')

    for(let numbEmotion = 0, i = 0; i < res.length; i++){

        if(res[i].emotion) {  
            addRecommendationIMG(recommendations[i], PHOTOS.photos[i])
            addEmotion(emotions[numbEmotion++], res[i].emotion)
            addTextRecommendation(recommendations[i], res[i].text)
        }   
        else{
            addNoEmotion(recommendations[i], PHOTOS.photos[i]);
        }
    }

    deleteLoading();
    document.querySelector('#block3').scrollIntoView({
        behavior: "smooth",
        block: "start"
    });
})
}


function wiewRecommendations(){
if(document.querySelector('#textSadness').classList.contains('open')) openTextSadness();
if(document.querySelector('#textHappiness').classList.contains('open')) openTextHappiness();
if(document.querySelector('#textAnger').classList.contains('open')) openTextAnger();
if(document.querySelector('#textDisgust').classList.contains('open')) openTextDisgust();
if(document.querySelector('#textFear').classList.contains('open')) openTextFear();
document.querySelector('.blockRecommendations').classList.toggle('open');
}


function newFuncAddText(){
let end = 0;
let timerId;
return function addText(target, text) {
    target.classList.toggle('open')
    if(!target.classList.contains('open')){
        if(timerId) clearInterval(timerId)
    }
    else{
        let i = end;
        timerId = setInterval(() => { 
            if(!text[i]) {clearInterval(timerId)} 
            else { 
                if(text[i] === "\n")  target.innerHTML += '<br>';  //  перевод строки

                target.innerHTML+=text[i]; 
                end = ++i;  
            }
        }, 30);
    }
}
}

let addtextSadness = newFuncAddText();
let addTextHappiness = newFuncAddText();
let addTextAnger = newFuncAddText();
let addTextDisgust = newFuncAddText();
let addTextFear = newFuncAddText();




let textSadness;
async function openTextSadness (){
if(!textSadness){
    let response = await fetch('http://localhost:8080/txt/textSadness.txt');
    textSadness = await response.text();
}
addtextSadness(document.querySelector('#textSadness'), textSadness);
}


let textHappiness;
async function openTextHappiness(){
if(!textHappiness){
    let response = await fetch('http://localhost:8080/txt/textHappiness.txt');
    textHappiness = await response.text();
}
addTextHappiness(document.querySelector('#textHappiness'), textHappiness);
}

let textAnger;
async function openTextAnger() {
if(!textAnger){
    let response = await fetch('http://localhost:8080/txt/textAnger.txt');
    textAnger = await response.text();
}
addTextAnger(document.querySelector('#textAnger'), textAnger);
}

let textDisgust;
async function openTextDisgust (){
if(!textDisgust){
    let response = await fetch('http://localhost:8080/txt/textDisgust.txt');
    textDisgust = await response.text();
}   
addTextDisgust(document.querySelector('#textDisgust'), textDisgust);
}

let textFear;
async function openTextFear (){
if(!textFear){
    let response = await fetch('http://localhost:8080/txt/textFear.txt');
    textFear = await response.text();
}   
addTextFear(document.querySelector('#textFear'), textFear);
}




function showText(target){
target.classList.add('open')
}


window.onload = () => {
const options = {
    root: null,  // набдюдатель(null = viewport)
    rootMargin: "0px",
    threshold: 0.3,
}

const observer = new IntersectionObserver((entries, observer) => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {         
            showText(entry.target)
        }

    })
}, options)

const arr = document.querySelectorAll('.textNeuralNetwork') // наблюдаемый
arr.forEach(i => {
    observer.observe(i)
})
}