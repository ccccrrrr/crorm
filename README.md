#   c r o r m  
  
 a n   i m i t a t i o n   o r m   f r o m   g o r m  
  
 c r o r m   s u p p o r t s   b a s i c   C R U D   o p e r a t i o n s   b a s e d   o n   m y s q l  
  
 w h i l e   n o t   l i k e   g o r m   y o u   o p e r a t e   m a i n l y   b y   g o r m . D B ,   c r o r m   d o   m o s t   o p e r a t i o n s   i n   c r o r m . T a b l e  
  
 # #   O p e n   a   d a t a b a s e  
  
 y o u   c a n   c o n f i g u r e   y o u r   m y s q l   c o n n e c t i o n   i n f o   u s i n g   D B C o n f i g   s t r u c t  
  
 ` ` ` g o  
 t y p e   D B C o n f i g   s t r u c t   {  
       U s e r N a m e           s t r i n g  
       U s e r P a s s w o r d   s t r i n g  
       P o r t                   s t r i n g  
       I P                       s t r i n g  
       D B N a m e               s t r i n g  
 }  
 / /   a f t e r   c o n f i g u r a t i o n ,   c o n n e c t  
 d b ,   e r r   : =   O p e n ( s t a n d a r d C o n f i g )  
 ` ` `  
  
 # #   S e l e c t   a   t a b l e  
  
 I n   c r o r m ,   y o u   c a n   c h o o s e   a   e x i s t e d   t a b l e ,   c r e a t e   a   n e w   t a b l e   o r   o v e r r i d e   a n   o l d   o n e .   a l l   o f   t h e   o p e r a t i o n   r e t u r n s   t h e   t a b l e   l i n k s ,   o r   r e t u r n s   t h e   e r r o r   i f   t h e r e   i s   t y p e   m i s m a t c h   o r   s o m e t h i n g   w r o n g   t o   g e t   t a b l e   i n f o r m a t i o n  
  
 ` ` ` g o  
 / /   p u r e   c r e a t e   t a b l e   o p e r a t i o n   f o l l o w e d   b y   t h e   t y p e   o f   s t a n d a r d   s t r u c t  
 t a b l e ,   e r r   : =   d b . C r e a t e T a b l e ( & s t a n d a r d S t r u c t )  
  
 / /   i f   t h e   t a b l e   a l r e a d y   e x i s t s ( h a v e   t h e   s a m e   n a m e ) ,   c r o r m   w i l l   j u s t   o v e r r i d e   i t  
 t a b l e ,   e r r   : =   d b . C r e a t e O r O v e r r i d e T a b l e ( & s t a n d a r d S t r u c t )  
  
 / /   c r o r m   w i l l   s e a r c h   f o r   t h e   e x i s t e d   t a b l e   i n   t h e   d a t a b a s e ,   a n d   r e t u r n s   t a b l e   i n f o r m a t i o n   i f   a l l   t h e   n a m e   a n d   t y p e   m a t c h e s ( f o l l o w i n g   a   v e r y   s t r i c t   r u l e )  
 t a b l e ,   e r r   : =   d b . S y n c ( & s t a n d a r d S t r u c t )  
 ` ` `  
  
 # #   I n s e r t  
  
 # #   D e l e t e  
  
 # #   F i n d   a n d   F i r s t  
  
 # #   U p d a t e  
  
  
  
 