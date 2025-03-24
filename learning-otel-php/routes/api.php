<?php

use App\Http\Controllers\TodoController;
use Illuminate\Support\Facades\Route;

Route::prefix('todo')
    ->controller(TodoController::class)
    ->group(function () {
        Route::get('/', 'index');
        Route::get('/error', 'causeError');
        Route::get('/{id}', 'show');
        Route::delete('/{id}', 'destroy');
        Route::post('/', 'store');
    });
