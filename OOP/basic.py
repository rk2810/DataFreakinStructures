# A cliche program that demonstrates classes in Python.

class Employee:
    """A class is just a blue print for creating instances."""
    def __init__(self, first_name, last_name, salary):
        self.first_name = first_name
        self.last_name = last_name
        self.slary = salary

    def fullname(self):
        return f'{self.first_name + " " + self.last_name}'


employee1 = Employee('Rohan', 'K', 50000)   # An instance of class Employee
print(employee1.first_name)
print(employee1.fullname())
