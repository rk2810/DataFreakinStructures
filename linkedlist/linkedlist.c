#include<stdio.h>
#include<stdlib.h>
// node declaration

struct Node{
	int data;
	struct Node *next;
};

// function to print nodes
void printList(struct Node *n){
	while(n!=NULL){
		printf("%d\n",n->data);
		n=n->next;
	}
}

// lets make 3 nodes
int main(){
	struct Node* head=NULL;
	struct Node* second=NULL;
	struct Node* third=NULL;

	// allocate 3 blank nodes

	head=(struct Node*)malloc(sizeof(struct Node));
	second=(struct Node*)malloc(sizeof(struct Node));
	third=(struct Node*)malloc(sizeof(struct Node));
// Assign head
	head->data=1234;
	head->next=second;
// Assign second
	second->data=5678;
	second->next=third;
// Assign third
	third->data=91011;
	third->next=NULL;
// Print the damn list.
printList(head);
return 0;
}