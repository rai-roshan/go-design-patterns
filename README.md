# go-design-patterns

## Decorator | [reading material](https://refactoring.guru/design-patterns/decorator) | [video material](https://www.youtube.com/watch?v=F365lY5ECGY&list=PLJbE2Yu2zumAKLbWO3E2vKXDlQ8LT_R28&index=1)

- Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.
- Using decorators you can wrap objects countless number of times since both target objects and decorators follow the same interface. The resulting object will get a stacking behavior of all wrappers.
- Its main components are:

1. The Component declares the common interface for both wrappers and wrapped objects. (eg: `FinanceService Interface`)
2. Concrete Component is a class of objects being wrapped. It defines the basic behavior, which can be altered by decorators. (eg: `financeService stucture instance/object`)
3. The Base Decorator class has a field for referencing a wrapped object. The field’s type should be declared as the component interface so it can contain both concrete components and decorators. The base decorator delegates all operations to the wrapped object. ( eg: the field name `next`, which carry the main/core component/instance of finance service)
4. Concrete Decorators define extra behaviors that can be added to components dynamically. Concrete decorators override methods of the base decorator and execute their behavior either before or after calling the parent method. (eg: `loggingMiddleware struct`)

## FACTORY | [reading material](https://refactoring.guru/design-patterns/factory-method) | [video material](https://www.youtube.com/watch?v=-1xgg7yUlUc&list=PLJbE2Yu2zumAKLbWO3E2vKXDlQ8LT_R28&index=3)
- Factory Method is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

1. The `Product` declares the interface, which is `common to all objects` that can be produced by the `creator` and its subclasses.
2. `Concrete Products` are different implementations of the product interface.
3. The Creator class declares the `factory method` that returns `new product objects`. It’s important that the return type of this method matches the product interface. You can declare the factory method as abstract to force all subclasses to implement their own versions of the method. As an alternative, the base factory method can return some default product type.
4. `Concrete Creators` override the base factory method so it returns a different type of product.
```
Note that the factory method doesn’t have to create new instances all the time. It can also return existing objects from a cache, an object pool, or another source.
```

## ADAPTER
