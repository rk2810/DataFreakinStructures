# Let's now focus on something called a classmethod and a staticmethod.
class Employee():
    """A class is just a blue print for creating instances."""

    num_of_emps = 0    # this is a class variable to keep track of number of employees.
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

    @classmethod
    def set_raise_amt(cls, amount):    # What's special about this is the argument where 'cls' is class that goes in as an argument.
        cls.raise_default = amount

    @classmethod
    def from_string(cls, data_string):    # let's tackle a case where user's input is just a ':' separated input for instantiation like 'Joe:Brown:50000'
        first, last, sal = data_string.split(':')
        return cls(first, last, sal)

    @staticmethod    # Difference between static and class method is.... well, class method takes class as argument and static method does not take any argument.
    def is_even(num):   # just a regular function that has nothin to do with class or instance.
        return 'Even.' if num % 2 == 0 else 'Odd.'

Employee.set_raise_amt(1.07)    # let's set raise amount to 7%

# Let's try using our from_string classmethod
employee3 = Employee.from_string('John:Doe:50000')

employee1 = Employee('Rohan', 'K', 50000)   # An instance of class Employee
employee2 = Employee('Test', '1', 60000)
print(employee1.first_name)
print(employee2.fullname())
print(employee3.fullname())
print(Employee.num_of_emps)     # returns number of instances/employees.
print(employee1.apply_raise())
employee2.raise_default = 1.05  # override default raise for employee2.
print(Employee.apply_raise(employee2))

# Magic method to print namespaces of class and instances in form of a dictionary.
print(Employee.__dict__)
print(employee1.__dict__)
print(employee2.__dict__)
print(Employee.is_even(5))

