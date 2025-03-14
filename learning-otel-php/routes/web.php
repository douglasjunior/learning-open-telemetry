<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\TodoController;

Route::prefix('api')->group(function () {

    Route::prefix('todo')
        ->controller(TodoController::class)
        ->group(function () {
            Route::get('/', 'index');
            Route::get('/{id}', 'get');
            Route::delete('/{id}', 'delete');
            Route::post('/', 'create');
        });

});
