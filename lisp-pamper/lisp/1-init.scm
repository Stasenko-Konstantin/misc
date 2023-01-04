#lang r5rs

(define env.init '())
(define empty-begin 813)
(define the-false-value (cons "false" "boolean"))

(define (extend vars vals env)
  (cond
    ((pair? vars)
     (if (pair? vals)
         (cons (cons (car vars) (car vals))
               (extend (cdr vars) (cdr vals) env))
         (wrong! "Too few values")))
    ((null? vars)
     (if (null? vals)
         env
         (wrong! "Too many values")))
    ((symbol? vars) (cons (cons vars vals) env))))

(define (eprogn exps env)
  (if (pair? exps)
      (if (pair? (cdr exps))
                 (begin (evaluate (car exps) env)
                        (eprogn (cdr exps) env))
                 (evaluate (car exps) env))
      empty-begin))

(define (evlis exps env)
  (if (pair? exps)
      (let ((arg1 (evaluate (car exps) env)))
        (cons arg1
            (evlis (cdr exps) env)))
      '()))

(define (lookup var env)
  (if (pair? env)
      (if (eq? (caar env) var)
          (cdar env)
          (lookup var (cdr env)))
      (wrong! "No such binding" var)))

(define (update! var val env)
  (if (pair? env)
      (if (eq? (caar env) var)
          (begin (set-cdr! (car env) val)
                 val)
          (update! var val (cdr env)))
      (wrong! "No such binding" var)))

(define (invoke fn args env)
  (if (procedure? fn)
      (fn args env)
      (wrong! "Not a function" fn)))

(define (make-function vars body env)
  (lambda (vals)
    (eprogn body (extend env vars vals))))

(define (wrong! . args)
  (let ((print (lambda (x)
                 (begin
                   (display x)
                   (display " ")))))
    (map print args))
  '())

(define (atom? exp)
  (not (or (null? exp)(pair? exp))))

(define (evaluate exp env)
  (if (atom? exp)
      (cond
        ((symbol? exp) (lookup exp env))
        ((or (number? exp) (string? exp) (char? exp)
             (boolean? exp) (vector? exp))
         exp)
        (else (wrong! "Cannot evaluate exp")))
      (case (car exp)
        ((quote)  (cadr exp))
        ((if)     (if (not (eq? (evaluate (cadr exp) env) the-false-value))
                      (evaluate (caddr exp) env)
                      (evaluate (cadddr exp) env)))
        ((begin)  (eprogn (cdr exp) env))
        ((set!)   (update! (cadr exp) (evaluate (caddr exp) env) env))
        ((lambda) (make-function (cadr exp) (cddr exp) env))
        (else     (invoke (evaluate (car exp) env)
                          (evlis (cdr exp) env) env)))))

;(evaluate (if #t 1 2) env.init)




