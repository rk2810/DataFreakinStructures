# Let's talk about class vs instance variables.

class Employee:
    """A class is just a blue print for creating instances."""

    num_of_emps = 0     # this is a class variable to keep track of number of employees.
    raise_default = 1.04    # Let's give our employees a raise they deserve.

    def __init__(self, first_name, last_name, salary):
        self.first_name = first_name
        self.last_name = last_name
        self.salary = salary

        Employee.num_of_emps += 1

    def fullname(self):
        return f'{self.first_name + " " + self.last_name}'

    def apply_raise(self):
        return (self.salary * self.raise_default)


employee1 = Employee('Rohan', 'K', 50000)   # An instance of class Employee
employee2 = Employee('Test', '1', 60000)
print(employee1.first_name)
print(employee2.fullname())
print(Employee.num_of_emps)     # returns number of instances/employees.
print(employee1.apply_raise())
employee2.raise_default = 1.05  # override default raise for employee2.
print(Employee.apply_raise(employee2))

# Magic method to print namespaces of class and instances in form of a dictionary.
print(Employee.__dict__)
print(employee1.__dict__)
print(employee2.__dict__)

