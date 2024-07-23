<?php

use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});
Route::get('home', function(){
    return view ('FE.index');
});

Route::get('properties', function(){
    return view ('FE.pages.properties');
});

Route::get('properti_detail', function(){
    return view ('FE.pages.properti_detail');
});

Route::get('contact', function(){
    return view ('FE.pages.contact');
});