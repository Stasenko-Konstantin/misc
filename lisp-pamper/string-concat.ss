(define (string-concat . strs)
  (apply string-append strs))

;; костыль:
(define (string-concat2 . strs)
  (cond
   ((null? strs) "")
   ((null? (cdr strs)) (car strs))
   (else (let*
      ([s1 (car strs)]         
	[s2 (cadr strs)]        
	[l1 (string-length s1)] 
	[l2 (string-length s2)] 
	[l3 (+ l1 l2)]          
	[s3 (make-string        
	     (+ l1 l2)
	     #\0)])
     (begin
       (string-copy! s1 0 s3 0 l1)
       (string-copy! s2 0 s3 l1 l2)
       (apply string-concat (cons s3 (cddr strs))))))))
