# SUMMARY

## 1. Clean Architecture (CA)

Clean Architecture is an architectural concept or solution from Robert C. Martin that we can use in the application development process. Good architecture is a separation of concerns using layers to build a modular, scalable, maintainable, and testable application. Because working on a programming project is similar to planning and building a residential district. If you know that the district will be expanding in the near future, you need to keep space for future improvements. Without that, you will be forced to destroy buildings and streets to make space for the new building.

## 2. Constraint & Benefit

The constraint before designing clean architecture are:

1. Independent of Framework
    
    This allows you to use such frameworks as tools, rather than having their limited constraint.
    
2. Testable
    
    The business rules can be tested without any other stuff.
    
3. Independent of UI
    
    The UI can change easily, without changing the rest of the system.
    
4. Independent of Database
    
    Your business rules are not bound to the database, you can swap out the database easily.
    
5. Independent of any external
    
    In fact, your business rules simply don’t know anything at all outside the world.
    

The benefit of implementing clean architecture are:

- A standard structure, so it’s easy to find your way into the project
- Faster development in the long term
- Mocking dependencies becomes trivial in unit tests
- Easy switching from prototypes to proper solutions (e.g. changing in-memory storage to a SQL Database)

## 3. CA Layer

- **Entities layer** (optional)
    
    Business objects as they reflect the concepts that your app manages
    
- **Use Case - Domain Layer**
    
    Contains business logic of your app
    
- **Controller - Presentation Layer**
    
    Present data to a screen and handle user interactions
    
- **Driver - Data layer**
    
    Manages application data e.g. retrieve data from the network, manage data cache