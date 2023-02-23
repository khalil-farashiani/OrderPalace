# OrderPalace

is a order management system designed. Built with the Clean Architecture in mind, OrderPalace is organized around a core domain model and features modular use cases, adapters, and frameworks that can be easily customized to meet the specific needs of each business.

The system features a RESTful API that allows orders to be submitted and processed asynchronously, with a Redis queue serving as the primary message broker. Orders are stored in a MySQL database, and can be queried and managed via a user-friendly web interface.