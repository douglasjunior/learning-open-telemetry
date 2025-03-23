## Requirements

- PHP 8.2 or higher
- Composer
- NodeJS and npm
- OpenTelemetry extension https://pecl.php.net/package/opentelemetry
- gRPC extension https://pecl.php.net/package/gRPC

## Try

1. Run `composer install`
1. Copy `.env.example` to `.env`
1. Run `php artisan migrate`
1. Run `composer run dev`
1. Open `http://localhost:8000/docs` in your browser
