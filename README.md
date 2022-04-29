# Tugas Besar 3 kelompok 47 Kobok DNA
> Pada tugas besar 3 ini mahasiswa diminta membuat web pendeteksi penyakit berdasarkan DNA menggunakan algoritma string matching KMP dan BM.

## Table of Contents
* [General Info](#general-information)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Screenshots](#screenshots)
* [Setup](#setup)
* [Usage](#usage)
* [Members](#members)

## General Information
- Tugas Besar kali ini memiliki topik string matching dengan algoritma yang lebih mangkus dibanding brute force.
- String matching ini akan digunakan untuk mencocokan DNA penyakit dengan DNA pengguna yang akan dicek
- Selain string matching, akan digunakan regex untuk mengsanitasi input file dari mengguna serta untuk melakukan query pada database.

## Technologies Used
- Vue 3
- Golang
- MongoDB

## Features
- Tes DNA(dengan input sequence DNA pengguna serta nama penyakit yang akan dicari)
- Menambah penyakit ke database(menerima input sequence penyakit)
- Mencari hasil tes yang pernah dilakukan dengan memasukan query


## Screenshots
![image](https://user-images.githubusercontent.com/74661051/165987656-3511b584-5439-41ab-8893-cd093546f3ac.png)


## Setup
- Install node js
- Install Vue
- Install golang

## Usage
1. Nyalakan front end dengan command ```npm run serve``` pada folder FrontEnd
2. Nyalakan BackEnd dengan command ```go run main``` pada folder BackEnd
3. Klik link yang muncul ketika menjalankan npm run serve(pilih salah satu)

## Members
- 13520028 - Timothy Stanley Setiawan
- 13520091 - Andreas Indra Kurniawan
- 13520137 - Adzka Ahmadetya Zaidan
