#lang racket

;; somewhat from htdp

(require 2htdp/image
         2htdp/universe)

(define ufo (overlay (rectangle 40 4 "solid" "blue")
                     (circle 10 "solid" "green")))

(define (compute-border scene-coord image-coord)
  (let* ([border 10]
         [edge (- scene-coord border)])
    (cond
      [(< image-coord border) border]
      [(> image-coord edge) edge]
      [#t image-coord])))

(define (my-place-image image scene-width scene-height image-width image-height)
  (place-image image
               (compute-border scene-width image-width)
               (compute-border scene-height image-height)
               (empty-scene scene-width scene-height)))

(define (distance time speed)
  (* time speed))

(animate (λ (time)
           (my-place-image ufo 300 200 150 (distance time 3))))
