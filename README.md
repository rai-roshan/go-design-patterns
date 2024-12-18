# go-design-patterns


## Decorator | [reading material](https://refactoring.guru/design-patterns/decorator) | [video material](https://www.youtube.com/watch?v=F365lY5ECGY&list=PLJbE2Yu2zumAKLbWO3E2vKXDlQ8LT_R28&index=1)
- Decorator is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.
- Using decorators you can wrap objects countless number of times since both target objects and decorators follow the same interface. The resulting object will get a stacking behavior of all wrappers.
- Its main components are:
1. The Component declares the common interface for both wrappers and wrapped objects. (eg: `FinanceService Interface`)
2. Concrete Component is a class of objects being wrapped. It defines the basic behavior, which can be altered by decorators. (eg: `financeService stucture instance/object`)
3. The Base Decorator class has a field for referencing a wrapped object. The field’s type should be declared as the component interface so it can contain both concrete components and decorators. The base decorator delegates all operations to the wrapped object. ( eg: the field name `next`, which carry the main/core component/instance of finance service)
4. Concrete Decorators define extra behaviors that can be added to components dynamically. Concrete decorators override methods of the base decorator and execute their behavior either before or after calling the parent method. (eg: `loggingMiddleware struct`)
