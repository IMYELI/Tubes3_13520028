<template>
    <img src="../assets/kobo2_resized.png" style="position:fixed;bottom:0;right:0">
    <div v-if="mainMenu" class = "box-wrapper">
        <h1 style="margin-bottom: 100px;">
            Main Menu
        </h1>
        <div class = "menu-option">
            <button @click="{ mainMenu = false; testDNA = true}" class = "menu-option-button">Test DNA</button>
            <button @click="{ mainMenu = false; addPenyakit = true}" class = "menu-option-button">Penambahan Penyakit</button>
            <button @click="{ mainMenu = false; searchDatabase = true}" class = "menu-option-button">Search Database</button>
        </div>
    </div>
    
    <div v-if="testDNA" @submit.prevent="compress" action="" class = "box-wrapper">
        <button @click="{mainMenu = true; testDNA = false; reset()}" class="back-button">back</button>
        <h1 style="margin-bottom: 20px;">
            Test DNA
        </h1>
        <div class="input-option-wrapper">
            <div class = "input-option">
                <h3>Nama Pengguna</h3>
                <input type="text" v-model="namaPengguna" class = "menu-option-input" placeholder="<pengguna>">
            </div>
            <div class = "input-option">
                <h3>Upload Sequence</h3>
                <label class = "custom-file-upload">
                    <input type="file" @change="onFileUploaded($event.target.files)" accept=".txt">
                    <a v-if="uploaded">
                        {{namaFile}}
                    </a>
                    <a v-else>
                    Upload Sequence...
                    </a>
                </label>
            </div>
            <div class = "input-option">
                <h3>Prediksi Penyakit</h3>
                <input type="text" v-model="namaPenyakit" class = "menu-option-input" placeholder="<penyakit>">
            </div>
            <div v-if="showResult">
                {{tanggal}} - {{namaPengguna}} - {{namaPenyakit}} - {{persentase}}% - {{diagnosis}}
            </div>
            <div v-if="showError" style="color:red">
                {{errorMessage}}
            </div>
        </div>
        <div class = "radio-button">
            <input type="radio" v-model="isKMP" value = "1">KMP
            <input type="radio" v-model="isKMP" value = "0">BM
        </div>
        <div class="submit-div">
            <button @click="tesDNA();showError = false;showResult = false">Submit</button>
        </div>
    </div>

    <div v-if="addPenyakit" @submit.prevent="compress" action="" class = "box-wrapper">
        <button @click="{mainMenu = true; addPenyakit = false; reset()}" class="back-button">back</button>
        <h1 style="margin-bottom: 20px;">
            Tambah Penyakit
        </h1>
        <div class="input-option-wrapper">
            <div class = "input-option">
                <h3>Nama Penyakit</h3>
                <input type="text" v-model="namaPengguna" class = "menu-option-input" placeholder="<penyakit>">
            </div>
            <div class = "input-option">
                <h3>Upload Sequence</h3>
                <label class = "custom-file-upload">
                    <input type="file" @change="onFileUploaded($event.target.files)" accept=".txt">
                    <a v-if="uploaded">
                        {{namaFile}}
                    </a>
                    <a v-else>
                        Upload Sequence...
                    </a>
                </label>
            </div>
        </div>
        
        <div class="submit-div">
            <button>Submit</button>
        </div>
    </div>


    <div v-if="searchDatabase" action="" class = "box-wrapper">
       <button @click="{mainMenu = true; searchDatabase = false; reset()}" class="back-button">back</button>
        <h1 style="margin-bottom: 20px;">
            Search Database
        </h1>
        <div class="input-option-wrapper">
            <input type="text" class="database-search-input" @keyup.enter="searchQuery()">
        </div>
        <div v-if="queryEntered">
            <ul>
                <li v-for="item in arrTest.slice(lowIndex,highIndex)" :key="item.index" class="hasil-query">
                    {{item.namaPenyakit}}
                </li>
            </ul>
            <div class="input-option-wrapper">
                <button class="prev-next-button" @click="prevQuery()">
                    &lt;
                </button>
                <button class="prev-next-button" @click="nextQuery()">
                    &gt;
                </button>
            </div>
        </div>
    </div>
</template>

<script>
// import axios from 'axios';
import axios from 'axios'

export default {
    
    data(){
        return{
            mainMenu: true,
            addPenyakit: false,
            testDNA: false,
            searchDatabase: false,
            uploaded: false,
            queryEntered: false,
            showResult: false,
            showError: false,
            diagnosis: "False",
            namaPengguna: "",
            namaFile: "",
            namaPenyakit: "",
            newPenyakit: "",
            respPenyakit: "",
            isiFile: "",
            tanggal: "",
            errorMessage: "",
            arrTest: [{index: 1,namaPenyakit: "a",namaFile: "a"},
                        {index: 2,namaPenyakit: "b",namaFile: "b"},
                        {index: 3,namaPenyakit: "c",namaFile: "c"},
                        {index: 4,namaPenyakit: "d",namaFile: "d"},
                        {index: 5,namaPenyakit: "e",namaFile: "e"},
                        {index: 6,namaPenyakit: "f",namaFile: "f"},
                        {index: 7,namaPenyakit: "g",namaFile: "g"},
                        {index: 8,namaPenyakit: "h",namaFile: "h"},
                        {index: 9,namaPenyakit: "i",namaFile: "i"},
                        {index: 10,namaPenyakit: "j",namaFile: "j"},
                        {index: 11,namaPenyakit: "k",namaFile: "k"},
                        {index: 12,namaPenyakit: "l",namaFile: "l"}
                        ],
            lowIndex: 0,
            highIndex: 5,
            queryLength: 12,
            dataPerPage: 5,
            persentase: 0,
            isKMP: 1,
        }
    },
    methods:{
        onFileUploaded(listFile){
            this.uploaded = true;
            this.namaFile = listFile[0].name;
            const fileReader = new FileReader();
            fileReader.onload = (res) =>{
                console.log(res.target.result);
                this.isiFile = res.target.result;
            };
            fileReader.readAsText(listFile[0]);
        },

        reset(){                   //Digunakan untuk mereset semua state dan server json
            this.namaPengguna = "";
            this.namaFile = "";
            this.namaPenyakit = "";
            this.uploaded = false;
            this.queryEntered = false;
            this.showResult = false;
            this.showError = false;
            this.diagnosis = "False";
            this.persentase = 0;
            this.isiFile = "";
            this.tanggal = "";
            this.isKMP = 1;
        },

        nextQuery(){
            if(this.highIndex + this.dataPerPage >this.queryLength && this.queryLength%this.dataPerPage !== 0){
                this.lowIndex = this.queryLength-this.queryLength%this.dataPerPage;
                this.highIndex = this.queryLength;
                
            }else if(this.highIndex + this.dataPerPage <= this.queryLength){
                this.lowIndex = this.lowIndex + this.dataPerPage;
                this.highIndex = this.highIndex + this.dataPerPage;
            }
        },

        prevQuery(){
            if(this.lowIndex - this.dataPerPage < 0){
                this.lowIndex = 0;
                this.highIndex = this.dataPerPage;
            }else if(this.highIndex === this.queryLength && this.queryLength%this.dataPerPage !== 0){
                var x;
                x = this.queryLength % this.dataPerPage;
                this.lowIndex = this.lowIndex - this.dataPerPage;
                this.highIndex = this.highIndex - x;
            }else{
                this.lowIndex = this.lowIndex - this.dataPerPage;
                this.highIndex = this.highIndex - this.dataPerPage;
            }
        },
        searchQuery(){
            this.queryEntered = true;
        },
        tesDNA() {
            var data_pass = {"NamaPengguna": this.namaPengguna, "SequenceDNA": this.isiFile, "NamaPenyakit": this.namaPenyakit, "StringMatching": this.isKMP};
            
             /*eslint-disable*/
            axios({ method: "POST", url: "http://localhost:8888/TestDNA", data: data_pass, headers: {"content-type": "text/plain" } }).then(result => { 
                console.log(result.data)
                // console.log(data_pass)
                if(result.data["errors"] != null){
                    throw new Error(result.data["errors"]);
                }
                this.tanggal = result.data["Tanggal"];
                this.diagnosis = result.data["Diagnosis"];
                this.persentase = result.data["SkorKesamaan"]
                // console.log("dontol")
                /*eslint-enable*/
                // this.response = result.data;
                // this.respPenyakit = result.data[""];
                /*eslint-disable*/
                // console.log(result.data) 
                /*eslint-enable*/
                this.showResult = true;
                }).catch( error => {
                    /*eslint-disable*/
                    console.log(error);
                    this.errorMessage = error;
                    this.showError = true;
                    /*eslint-enable*/
        })},
        fileSelected(name, listFile){
            this.fileUpload = listFile[0]
            this.namafile = listFile[0].name
            const fileReader = new FileReader()
            fileReader.readAsDataURL(this.fileUpload)
            console.log(this.percentage)                    
            fileReader.addEventListener('load', ()=>{                   
                this.imageURL = fileReader.result
                this.test = {'base64' : this.imageURL, 'percentage' : this.percentage,'namaFile' : this.fileUpload.name} //Membuat data yang akan di send ke json
                fetch('http://localhost:3000/image',{
                    method:'POST', //Upload Sequence ke server json
                    headers:{'Content-Type':'application/json'},
                    body: JSON.stringify(this.test)
                })
            })
            
            

        },
        compress(){             //Fungsi yang dipanggil ketika tombol compress ditekan
            // this.statusUpload = CONVERTING_STATUS
            // this.startTimer()
            // axios.get(this.pathFlask,{ responseType: 'blob'})       //Melakukan request ke server flask
            //     .then((res)=>{
            //         this.stopTimer()
            //         this.time = this.time /1000
            //         this.fileGet = res.data
            //         console.log(this.fileGet)
            //         const fileReader = new FileReader()
            //         fileReader.readAsDataURL(this.fileGet)
            //          fileReader.addEventListener('load', ()=>{
            //              this.imageURL2 = fileReader.result
            //              console.log(this.imageURL2)
            //          })
                    
            //         this.statusUpload = SUCCESS_STATUS
            //     })
            //     .catch(err=>{console.log(err.message),this.reset()})

        },
        download(){             //Fungsi yang digunakan ketika tombol download ditekan
            var fileURL = window.URL.createObjectURL(this.fileGet)
            var fileLink = document.createElement('a')
            fileLink.href = fileURL
            this.namafile2 = this.namafile.split('.').slice(0,-1).join('.')    //Mengambil namafile
            this.ext = this.namafile.substring(this.namafile.lastIndexOf('.') + 1) //Mengambil extension file
            this.namafile2 = this.namafile2 + '_Compressed.'+this.ext   //Menambahkan compressed di belakang nama file
            fileLink.setAttribute('download',this.namafile2)
            document.body.appendChild(fileLink)
            fileLink.click()
        },
        startTimer(){
            this.time = 0           //Memulai timer
            this.timer = Date.now()
        },
        stopTimer(){
            this.time = Date.now()-this.timer       //memberhentikan timer
            clearInterval(this.timer)
        }
        
        
    },
    mounted(){
        // fetch('http://localhost:3000/image')        //Load data pada server json saat pertama kali website dinyalakan
        //     .then(res=>res.json())
        //     .then(data=>this.test = data)
        //     .catch(err=>console.log(err.message))
        // this.reset()
    },
    

}
</script>

<style>
    h1 {
        text-align: center;
        font-size: 50px;
        font-family: 'Roboto', sans-serif;
        color: rgb(121, 195, 255);
        margin-top: 10px;
        

    }
    body {
        background-image: url('../assets/cloud3.png');
    }
    .box-wrapper{
        width: 60%;
        margin: 20px auto; 
        min-height: 400px;
        min-width: 800px;
        background: white;
        text-align: justify;
        padding: 40px;
        border-radius: 20px;
    }
    .menu-option{
        text-align: center;
    }
    .menu-option button{
        border: 3px solid black;
        background-color: white;
        color: rgb(129, 162, 255);
        padding: 14px 28px;
        font-size: 16px;
        cursor: pointer;
        border-radius: 8px;
        margin: 0px 10px;
        min-width: 240px;
        min-height: 75px;
        border-color: #76c1ff;
        font-family: 'Times New Roman', Times, serif;
        font-size: 25px;
        -webkit-text-stroke-width: 0.2px;
        -webkit-text-stroke-color: black;
    }
    .menu-option-button:hover{
        background-color: #76c1ff;
        color: white;
        -webkit-text-stroke-width: 0px;
    }

    .input-option-wrapper{
        text-align: center;
    }

    .input-option{
        display: inline-block;
        margin: 20px;
    }

    .input-option input{
        display: block;
    }
    .input-option h3{
        /* margin-left: 5px; */
        text-align: left;
    }

    .back-button{
        position:absolute;
        border: 3px solid #76c1ff;
        background-color: white;
        color: rgb(129, 162, 255);
        padding: 14px 28px;
        font-size: 16px;
        cursor: pointer;
        border-radius: 8px;
        margin: 0px 10px;
        font-family: 'Times New Roman', Times, serif;
        font-size: 25px;
        -webkit-text-stroke-width: 0.2px;
        -webkit-text-stroke-color: black;
    }
    .back-button:hover{
        background-color: #76c1ff;
        color: white;
        -webkit-text-stroke-width: 0px;
    }

    input[type="file"] {
        display: none;
    }
    input[type="text"] {
        height: 25px;
    }

    .custom-file-upload {
        border: 1px solid #ccc;
        display: inline-block;
        padding: 6px 48px;
        cursor: pointer;
    }

    .database-search-input {
        height: 50px;
        width: 300px;
        max-width: 300px;
        max-height: 50px;
        padding: 0px 15px;
        font-family: 'Times New Roman', Times, serif;
        font-size: 25px;
    }
    
    .prev-next-button{
        background-color: #76c1ff;
        color: white;
        border-radius: 10px;
        font-size: 25px;
        border: 1px solid black;
    }

    .hasil-query{
        list-style: none;
        padding: 5px;
        border: 2px solid black;
        border-radius: 7px;
        margin: 10px 0px;
    }

    .submit-div{
        text-align:center;
    }

    .submit-div button{
        border: 3px solid #76c1ff;
        background-color: white;
        color: rgb(129, 162, 255);
        padding: 14px 28px;
        font-size: 16px;
        cursor: pointer;
        border-radius: 8px;
        margin: 20px 10px;
        font-family: 'Times New Roman', Times, serif;
        font-size: 25px;
        -webkit-text-stroke-width: 0.2px;
        -webkit-text-stroke-color: black;
    }

    .submit-div button:hover{
        background-color: #76c1ff;
        color: white;
        -webkit-text-stroke-width: 0px;
    }

    .radio-button{
        padding-right: 10px;
        text-align: center;
    }
</style>




     