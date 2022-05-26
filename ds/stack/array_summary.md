hi LittleWhiteBear

Today we learned how to implement funtions for an array:
- an array should have attributes like below:
    - length
    - data contained

- useful funtions of an array should have
    - Push(): add an item to the end of the array
        - increase the leng by 1
        - attach the new item to the end of the array using append from lib
    - Pop(): remove the last item of the array
        - decrease the leng by 2
        - used [:leng-1] to not include the last element
            example: in total 5 elements[1,2,3,4,5]
            length is 5
            [:data.length-1] the right number is the destination of truncation, included
    - Delete(index int): remove the element at the specified index
        - use the pop function
        - O(n) as it moves the index of everything in worst case scenario: delete the first element



Reverse string:
GOTCHAs:
1. validate input