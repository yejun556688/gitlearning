GOOF----LE-8-2.0©&      ] K 4   h­      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-blank-slot¤	g  add-carriage-return-slot¤	g  deal-cards-face-up¤											
															 ¤	g  give-status-message¤	g  new-game¤	g  set-statusbar-message¤	g  get-stock-no-string¤	g  string-append¤	g  _¤	f  Stock left:¤	f   ¤	g  number->string¤	 g  length¤	!g  	get-cards¤	"g  empty-slot?¤	#g  	get-value¤	$g  get-top-card¤	%g  king¤	&g  button-pressed¤	'g  
fill-it-up¤	(g  member¤	)	ÿ
 ¤	*g  
droppable?¤	+g  add-to-score!¤	,g  remove-card¤	-g  min¤	.g  button-released¤	/g  button-clicked¤	0g  button-double-clicked¤	1g  game-won¤	2g  get-hint¤	3g  	game-over¤	4g  club¤	5f  Remove the king of clubs.¤	6g  diamond¤	7f  Remove the king of diamonds.¤	8g  heart¤	9f  Remove the king of hearts.¤	:g  spade¤	;f  Remove the king of spades.¤	<g  hint-remove-king¤	=g  
hint-click¤	>g  get-suit¤	?g  
king-check¤	@g  	hint-move¤	Ag  horizontal-check¤	Bg  vertical-check¤	Cg  backslash-check¤	Dg  slash-check¤	Eg  get-options¤	Fg  apply-options¤	Gg  timeout¤	Hg  set-features¤	Ig  droppable-feature¤	Jg  
set-lambda¤C 5      h    ] 4	 >  "  G       hø  º  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4	

>  "  G  4>   "  G  		 C²      g  filenamef  neighbor.scm
	
							#			3			C			U			e			h			m			v			y			~		 		 		 		 		 		  		 ©		 ¬		 ±		 º		 Ê	 	 Ú	!	 ê	"	 í	"	 ò	"	 û	#	 þ	#		#		$		$		$		%	 	%	%	%	.	&	1	&	6	&	?	'	O	(	_	)	o	*	r	*	w	*		+		+		+		,		,		,	¢	-	¥	-	ª	-	³	.	¶	.	»	.	Ä	/	Ô	0	ä	1	ô	2	÷	2	ü	2		3		3		3		4		4		4	'	5	*	5	/	5	8	6	;	6	@	6	I	7	Y	8	i	9	y	:	|	:		:		;		;		;		<		<	£	<	¬	=	¯	=	´	=	½	>	À	>	Å	>	Î	?	Ó	?	Ø	?	á	B	÷	D	 d	ø
  g  nameg  new-game CR   h   n   ] 45 6     f       g  filenamef  neighbor.scm
	F
		G			G	 		
  g  nameg  give-status-message CR !     h    ®   ] 45444
5556 ¦       g  filenamef  neighbor.scm
	I
		J				J			J			J	"		K			K	!		K	)		K	!		K			J	 		
  g  nameg  get-stock-no-string CR"#$%   h0   Î   ] 
$  C4 5$  C44 55C      Æ       g  slot-id
		* g  	card-list		*  g  filenamef  neighbor.scm
	M
		N		
	N			O			N			P			P		%	P		(	P		)	P	 		*	  g  nameg  button-pressed C&R"'     h   U  ] 	$  84
5$  "  4
  >  "  G   	$  	 6C4 5$  	 64   >  "  G   	$  	 6C     M      g  slot-id
	  g  spaces	   g  filenamef  neighbor.scm
	R
		S		
	S				S			U			U			V		%	V	3	*	V		7	X		;	X		@	Y		D	Y		G	[		N	[		P	[		T	S		[	b		]	b			^	]		e	]		j	]	@	o	]		|	^	 	^	 	_	 	_	 	 	  g  nameg  
fill-it-up C'R"#$()  hx   »  ] 		 		 $  C
$  C45$  C	454455$  45$  6CC ³      g  
start-slot
		w g  	card-list		w g  end-slot			w g  dx			w g  dy			w  g  filenamef  neighbor.scm
	d
		e			e			f			f			e			g			g			h			h			g			e		%	i		)	i		/	j		3	i		6	k		@	i		E	l		J	l	 	L	l		M	m		P	m	 	X	m		Y	l		Z	l		^	i		_	n		e	n		g	n		k	i		q	o		s	o	 "		w	  g  nameg  
droppable? C*R*+,'-    h@   à   ]4 5$  (4	5$  45$  4 56CCC     Ø       g  
start-slot
		; g  	card-list		; g  end-slot			;  g  filenamef  neighbor.scm
	q
		r			r			s			r			t		'	r		*	u		5	u	 
		;	  g  nameg  button-released C.R"#$%+,'   hX   ¿   ]4 5$  "  44 55$  *4>  "  G  4 >  "  G   6C       ·       g  slot-id
		Q  g  filenamef  neighbor.scm
	w
		x			x			y			y			y		!	y		%	x		&	{		7	|		O	}	 		Q  g  nameg  button-clicked C/R    h   t   ]C    l       g  slot
		  g  filenamef  neighbor.scm
 
 		  g  nameg  button-double-clicked C0R"   h   ]   ] 6 U       g  filenamef  neighbor.scm
 
	 	 		
  g  nameg  game-won C1R12      h(   y   ] 4>   "  G  45 $  C6        q       g  filenamef  neighbor.scm
 
	 		 		 		! 	 		!
  g  nameg  	game-over C3R456789:; 
     h@   Û   ] &  6 &  6 &  6 &  	6C    Ó       g  suit
		<  g  filenamef  neighbor.scm
 
	
 		 		 		 		 		 		& 		* 		, 		4 		8 		: 	 		<  g  nameg  hint-remove-king C<R%#$=<>"? 	     hP     ]44 55$   444 5556 	$  4 5$  C 6C          g  slot-id
		L  g  filenamef  neighbor.scm
 
	 		 		 		 			 		 		 	/	  	9	( 	/	* 		, 			1 		5 		6 		; 	 	= 		A 			H 		J 		 		L  g  nameg  
king-check C?R#$@"A        h`   J  ]
 	$  "  	44 5544 55$  
  6 	$  4	 5$  C 6C B      g  slot-id
		_  g  filenamef  neighbor.scm
 
	 			 		 			 		 	"	  		! 		$ 	"	) 	0	+ 	"	- 		. 		/ 		3 		; 		= 			B 		F 		G 		N 	 	P 		T 			[  		]  		 		_  g  nameg  horizontal-check CAR#$@"B     hX   -  ]	44 5544	 55$   	 6 	$  4	 5$  C 6C     %      g  slot-id
		S  g  filenamef  neighbor.scm
 £
	 ¤		 ¤		 ¤		 ¥		 ¥		 ¥	+	 ¥		 ¥		  ¤		! ¤			% ¤		/ ¦		1 ¦			6 §		: ¤		; ¨		B ¨	 	D ¨		H §			O ©		Q ©		 		S  g  nameg  vertical-check CBR#$@"C  hh   I  ]
 	$  "  	44 5544	 55$   	 64	 5$  "   	$   6C  A      g  slot-id
		f  g  filenamef  neighbor.scm
 ¬
	 ­			 ­		 ­			 ®		 ®	"	  ®		! ¯		$ ¯	"	+ ¯	0	- ¯	"	/ ¯		0 ®		1 ®		5 ­		? °		A °			B ±		I ±	 	K ±		O ±			Y ²		] ­		b ³		d ³		 		f  g  nameg  backslash-check CCR#$@"D      hh   E  ] 	$  "  	44 5544	 55$   	 6 	$  4	 5$  C 6C     =      g  slot-id
		c  g  filenamef  neighbor.scm
 ¶
	 ·			 ·		 ·			 ¸		 ¸	"	  ¸		! ¹		$ ¹	"	+ ¹	0	- ¹	"	/ ¹		0 ¸		1 ¸		5 ·		? º		A º			F »		J ·		K ¼		R ¼	 	T ¼		X »			_ ½		a ½		 		c  g  nameg  slash-check CDR?"ABCD        h   .  ]45  $   C4	5$  "  45  $   C4	5$  "  45  $   C4	5$  "  45  $   C4	5$  C	6       &      g  t
	  g  t
	)  g  t
	J  g  t
	k   g  filenamef  neighbor.scm
 À
	 Á		 Á		 Â		 Â		$ Ã		) Á		5 Ä		? Ä		E Å		J Á		V Æ		` Æ		f Ç		k Á		w È	  È	  É	 	 
  g  nameg  get-hint C2R     h   W   ] C    O       g  filenamef  neighbor.scm
 Ë
 		
  g  nameg  get-options CER    h   o   ]C    g       g  options
		  g  filenamef  neighbor.scm
 Í
 		  g  nameg  apply-options CFR    h   S   ] C    K       g  filenamef  neighbor.scm
 Ï
 		
  g  nameg  timeout CGR4HiIi>  "  G  Jii&i.i/i0i3i1i2iEiFiGi*i6      ÿ       g  filenamef  neighbor.scm		
õ	
	F
y	I
		M
	R
Þ	d
	q
R	w
ß 
X 
 
V 
â 
­ 
P £
 ¬
è ¶
É À
: Ë
Â Í
. Ï
/ Ñ
z Ó
 	z
   C6 