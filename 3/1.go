package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var input = "(%%from() when()mul(73,623)when()mul(793,458)'~where()how()?how(569,237)/[mul(709,198)mul(395,622)$!what()select()^@/what()+mul(970,343)mul(75,7)^))mul(61,40)select()~why())'>where()%+mul(892,307),!(mul(412,807):*&what()+^why()<^why()mul(706,931)'{who())why()^?mul(953,62)(mul(461,410)when()don't()>^@<%who()'mul(365,15)(<^<# (where()mul(802,710)why()*[(  where()where()mul(684,352)&)&what()^<[>mul(246,913)+select()?how(489,271)when() }why(627,30);don't()+&<@where()when()mul(636,990)]/mul(767,759):mul(328,474)([^,-select()?(mul(825,353)select()where()*where()what()*-@mul(765,991)where(786,744)'--where()mul(990,947) mul(547,706)from()?mul(229,193)where(617,453)+@&where()]@%}/mul(128,550)<%mul(3,636)don't()^(+#>+mul(66,503) select()#when()%from()mul(59,150)how(),when()^:mul(614,438)where()*<[where();(mul(434,344) /how()$ (%?~when()mul(659,534)mul(809,367) where(42,397),);? ++select()mul(858,771)<*mul(106,962)^>@;# @?,*mul(208,462)'',)mul(762,748)+[?}>][^mul(126,384))@]): /(('mul(966,704)who()what()~%*]from()mul(825,633)~$)+ &mul(634,698)(how()<@why(102,647)mul(661,112)]&<(%&mul(25,649)who()select();mul(267,405)why(356,766)why()(where(),select()^)mul(552,557){!from()^&>>+mul(493,578)select(742,239):how(){^mul(836,239)who()/)-mul(259,726)how()-;where()[#@from()~mul(495,301);>[;mul(478,953)$#[*why(),{mul(774,653)how()^mul(469,614)!what();-->>;mul(369,74)$who()who()mul(311,382)< ![>?$mul(909,70)!+$how(257,485)<mul(278,404)] where()':mul(824,974)when(),;&@-?~>mul(377,363)where()&why()-/:(mul(285,466)where()(what()why()[where()&who()mul(701,477)(where(),-why(){mul(624,21)where()[,why()-!?+mul(937,219) mul(604,90)how()(')]why()/mul(627,697)*;what()/$select()do()~<from()%-^mul(605,52)>@%!&select()>['mul(597,962)why()where()mul(903,469); &,do()when(937,722)((~why();^mul(588,272);}?[mul(295,621)from(645,893)>]>#mul(900,24)>,;mul(574,932):}(-do()from()-+<'$how()mul(694where())select()){how()how() (mul(350,308)/{)}why()what(987,719)}!#don't();<how()[from()/-'mul(438,672)}!@who() }?)mul(952,413#>:}-: 'when()select()mul(547,749)$#mul(869,866)@{mul(334,736';'select(330,146)[~ :>mul(511,40)(mul(11,3 mul(547,132)[!>select();who()mul(424,102)*!~mul(725,298)~:{^?%*mul(514,355%'],,: -)mul(116,719)select()?/@;when()<who();!mul(352,211)+#$;{$>]mul(820,414)&when(),!?[:?don't()when()?>}+#how()mul(159$$},:]why()-mul(689,30)@^mul(310,593)where(800,717)^(/*/! from()<mul(7,727)select()&mul(259,310)where()^select()#@how(49,595)what(){>]mul(429,841)(-+#[mul(579,668)>what()&+,mul(850,283)don't()#?mul(758,673)^;)select(856,890)%:how(550,676)<,)mul(667,314)>:[;+mul(760,374)select();}*select()#how()$*mul(347,822)who()how(){mul(497,700):where(){what()mul(851,789)^:}when()why()where()(:mul(991,536)-mul(711,63)',(  who()]mul(798,573)>mul(190,153)$]from()!mul(592,256)-$-]mul(734?mul(482,742)#;mul(939,69)how()([how()<what()$$who()why()mul(149,831))/'when()mul(152,123)+* ','$where()mul(774,252) ;@?,;don't()how()>(mul(933,652)}mul(882,656),?how()#%}do()!select()>}>mul(981,750)mul(927,646)!what()$mul(380why()$]/[*?+)mul(67,435)>select()@-mul(819,795)?/? how())mul(215,234){,#/>how(513,708)'~>do()#}^^mul#how()how()how()+when():mul(992,63)mul(526,962)where()}who()who()+&'<,mul(730,728)}/,select()what()*who())[mul(687,974)-<:mul(259,420)\nwho()(&?>'+:?mul(483,827)*!what()[!/*mul(368,168)!&;^)?mul(629,217)]^{what()!,who()mul(83,255)^)mul(500,689)!-mul(592,556)%select()}from()-+mul(946*who()?,from()%/'~from(619,712){mul(747,249)/@+%when())mul(762,891)what(),when()?,#<:[mul(873,69)+ :what()+how()&[}who()mul(591,81){when()mul(151,432)@<what()-@]do()/mul(233,269)%&![+where()what()-why()mul(402,497):@from()where()':(how()mul(520,79)!]^$<select()when(478,105)mul(340,948),(mul(410,461)select()from()]-(@who()who())]mul(511,538)!who(6,403)do()<(where()what()mul(443,727)when()where()?&where()~mul(302,402)&mul(599,23))[how()[+!{mul(823,16)#%#]+~who(942,742))how()mul(502,890)^;<where()from()!~;+;mul*$<;mul(25,421)/:where() {^who()mul(652,45)#?from()^~; ,<mul(783,401)?who()what()mul(828,237)!]^when(126,561)mul(724,536){<*?don't()])why()/mul(974,752)%]$})mul(688,4)!'#:'@$mul(115,891)['where()when()<where()#@mul(391,949)what()>,$mul(459,691)#>';where()mul(537,593):/]-'#),+mul(900,640)%*{why()?from()mul(750,181)]{%;]mul(771,902)<>$]how());,mul(646,610)why()mul(644,958)[/{<*)mul(490,14)/where()why()::where()${+?mul(524,278)when()mul(609,995){*(,how()(where(18,555)-who()mul(456,815)do()[$:~(}@mul(987,108)<select()who(),>}@>when()<mul(561,928)}$how(984,377)+:what(){ %]mul(242,236)mul(401,270)mul(144@who(){from()-[mul(850,190)& #where()@^/'>mul(511,942)select()$(/<mul(404,904)<^when()'how()[when()$}}mul(891,878)mul(381,288)<+[[<from()who()-mul(282,737),who(){what()-!>mul(60,709)(mul(291?}where(670,848)who()@from()',mul(890,902)[){,&%}%mul(539,945)%$~what()do()#what()how(931,606)!,',what()mul(175,630)]@-&when(98,693)$when()when()mul(850,996~;:@who()<^}/'mul(103,549)how(482,381);when()mul(408,760)where()!mul(796,714)&^]}mul(557,352)how()+{+]select()mul(135,816)?##@<mul(812,707)(^^what()%&$]mul(305,996) ?what()[mul(19,420)~{;mul(451<who()mul(678,132)!:]select();'select(414,538)-what()@mul(775,858)+'{ :[mul(151,737)*mul(305,447)?+<mul(773,348)*% ~}!>+do()@;from()%(mul(571,910)don't()?)%^+%mul(738,321)who()>&-?what()don't(){[?why()who()mul(558,221)}${!'(+*who()do())how();#<>?<mul(824where()mul(351,367)#/ :how()mul(212,77)}<]]mul(769,809)#when(303,117)select()(<':!%(mul(126,148)]why()%[^why()where()]mul(288,573)~@]>%'why()>)mul(850,160);~[select() why()why()~;mul(276,634) +{mul(615,507):when()++#:mul(817,325??(mul(748,223)^<!&^] &/mul(755,745)&:!&who()]-#;,mul(674,378)why()&}-how(667,459)( mul(688,272)'<who()~do()mul(608,206)~$^:'*)?*mul(47,323)mul(612,590)({how()mul(966,673)'(^/from()where()who()mul(802,18)mul(8,840))how()'^#[}(mul(588,513)(<[;when())) how()don't())who(97,403)mul(642,367)where()when()who(),why()from()'mul(333,247)+[where()-]+*mul(7,926)^!where()%'#where(109,801)mul(247,920)+}-&]select()select(878,420)#mul'{::what()how() mul(281,809)~/,from()/+^how()@>mul(67,325)mul(10,782)@-~#]$when()'mul(750,235)$from()]mul(434,941)when()~mul(375,426)mul(804,124)from())^when()from()[mul(478,817)where()where()$what()from()why()what()mul%%??:&mul(948,781)mul(135,972)%}$mul(466,336)#]mul(856,500);~,'%{?#?&mul(586,708)why()how()who(145,202)*who()why()%'who()#mul(784,992){{#mul(829,121)&/select()@%])mul(775,997)[@#mul(419,536)when()mul(526,395)<$;from()  when()who()%!mul(631,345),why()%%~~-;-,mul(197,766){don't()%+}}}>,*select()mul(607,892)\nmul(213,399)mul(50,339)]what()when()mul(15,850) $when()@)<why()}mul(863,448)[select(324,8);[mul(258,387)[~)-)why()mul(385,665):];(select()mul(260,974)when()mul(433,796)what()@!!*{where()mul(281,280){what()%{mul(996,146)'how(137,696)why()@*$mul(18,665)#who()@/(*?how()select()what()mul(817,632)<>]how(967,387)*!mul(287,345) ?why():,when(70,954)/;mul(593,904)$}/!mul(892,400)where()where()where(){where()^~how()mul(383,495)mul(268,214)why()$:^]^mul(269,261):-%why()%mul(530]when()(?]^!/what()@do()select()!]mul(186,765)&select()?:why()why()how()why()mul(77,756)from(){),;from()'when()where()mul(199,620)%'::mul(119,703)'#from()from()!'mul(315,665)[from()'',from()}( /mul(771,150)# mul(596,340) )>^don't(),'~:who()where()-$mul(697,969)!}why()mul(21,213)}]<#'+where()]>mul(857,557)where()when()[>-:mul(149,573):{where()'from():mul(810,346))}why(665,355)where()^mul(656,308)what()how()] @mul(528,812)mul(679,912)mul(748,249)*+how()mul(268,67)&]+< mul(817,699)-$-who()%mul(96,615)<:~what(941,999)--mul(217,695)]$))when()#/<from()^mul(687,294)#<}why() mul(870[#:*:where(659,602);[^@mul(136,411)where()what()[~where()+don't()@~who()!when()<why()mul(374,239)mul(65,349)&mul(149,46)]/>+ how()~}mul(909,65){:<'mul(349,935)select()@select()<mul(290,106)mul(874,477)*}, &,?}#mul(18,640)mul(973,307)?@$ [}mul(340,58)()$mul(878,389)~>#mul(472,182)?@[^select()what()mul(9,773)#'how(50,359){how()mul(512,560)mul(424,565)who():*{mul?%{'>-<-mul(732,974)]what()?-mul(549,212),{mul(45,855)(*<()>mul(811,328),what()<>$select()),?mul(785,169)^,mul(271,507)mul(135,488)*when(529,46) <where()what()from() how()mul(11,833)mul)$]who()>#}*select()from()mul(662,175)#$(;why()['why(){/mul(472,524)}<}~-@]mul(561,924)when()@%select(803,251):/})mul(602,358)-!%%$]#who() mul(986who()}{how();$'how()mul(423,659)where() ;?{[][)mul(682,685)><#>):mul(567,299)?+what()<mul(766,452),?~![mul(18,685)(::@-;->!mul(278,215)^~what()':(who();{mul(769,517)~how(365,364)from()'}why()mul(707,262)/!/how()><mul(339,279);/$do()^mul(256,483)who()<!mul(553,945)mul(500how(448,619)select(){~?)!(mul(607,566)$mul(777,986)how(): 'who()mul(205,345)from();+#when():do()~(when()what())what()/mul(174,64)]how()mul(537,63){^;who())!<what()(mul-why()mul(699,680)'/}*mul(465,40)mul(353,270)<!?; -?mul(350,912) {](what(){,from()@mul(813,789)*$;how()select()];]mul(198,144),:,-why(),mul(862,852)*mul(965,644)select()mul(884,219)~@!mul(435,862)select()$why()how()^mul(165,208)mul(979,195);mul(164,454)where()@?mul(957,86))++@-% what(543,264)who()mul(372,43)*$<who()?}<mul(176,225)/mul(764,915):~what(),!!mul(767,820);%where()^]mul(637,588)mul(7,937)?when()mul(930,146)where()?mul(148,167))select()+how(70,96)do()from(){]-{mul(791,513)>what()}>:]]select()]mul(56,217)[!,when()'why()where()<mul(762,753)what()<'*{+when() mul(654,921))mul(724,639):&{@why(521,545)(?how()who()mul(66,494)&'where()mul(308,521) -%$]:](>mul(974,552)'[]'mul*,do()#,>}mul(765,313)](,>who()#!/usr/bin/perl +how()#;who())>@%mul(868,912)where() #) @'/mul(764*<{do(),,,'+^mul(173,253)\nfrom()[^!mul(447,170)],/why()from()#mul(390,953)when()who()<>)%'@mul(830,176)what()!+:@mul))#>mul(350,913):,)*+mul(341,997)#+ mul(49,490):where()})select()when()mul(434,911)do() from()&[,who()how()&mul(277,83)@select()^{!,what()>:mul(896,117)why()-/mul(22,967)$^!)~)mul(540,843)^>}@^*/~[select()mul(305,728)mul(234,965)where()&;>;<)mul(536,159):#when(907,337)~/[^ where()mul(255,249)mul(47,296)*'select())+don't()+how()why()^ *@select();mul(985,160)}-<why()mul(283,5)how()why()^from()mul(930,617)>{what()),<^-?$mul(506,985)select()'>do(){how()  from()why()mul(329,225)$mul<don't()*{(who()*+mul(603,836)@mul(947,431)who(638,671)!don't():,what() {%*when()~(mul(179,618)*why()&;>;*+%when()mul~]who()$~from()when()?)mul(833,539)/%;!%^~#mul(37,690)@<':*>#~(mul(527,38)}who()~how()mul(812,603)$^who()why()&+]*mul(743,253)<mul(164,237)from()!&+$ ~}#^mul(250,473):who()(how(),,?+^mul(91,450)&select()from()'#where()from()-{$mul(245,112);&why()why()where()'/mul(332,324)%,$:where()~who()when(896,261)when()mul(551,82)!]mul(316,126),$from();mul(164,359)?[/from() $$mul(384,892)[where()$how()mul(721,238)-&don't())when(177,679)why()from()^from()+when()mul(404,530):do()?from()#&(({;^%mul(378,543)-mul(126,360)mul(699,438)mul(336,574)why()where()>}*,+'-from(486,483)mul(811,380)when()how()from()from()what()how()!mul(7,69);~!]^){why()mul(478,680):!/mul(227,472)why()from()]&mul(376,566):)! what()mul(241,708)select()]mul(410,202)what()where())do(){^mul(560,185);!<@mul(711,222,:@$*~*mul(103,347)@from())from()mul(657,475)/!;don't()!]who():$why()-}}:mul(451,895)^mul(981,444)-&-why()!#mul(100,601)where(){,how()<mul(213,723)mul(685,739[(mul(182,236),why(166,364)<@+*}&,mul'<mul(807,707)@)),(^do():;&mul(387,217)mul(878,247)mul(245,166) ?why()():/-+,do()! when()[~why(703,68)from()mul(694,278),: how()$*:[?mul(475,535-<<when(277,526),{/;what()?mul(798,822)^)^mul(138,477)how()]mul(663,377)why(568,281)how()[what()*<what()-@from()mul(655,71)<&/mul(35,807)select()]!,$&?,where()mul,/!}])/'mul(548,258)*)(;/why()don't(){>:[,*<[ ;mul(246,732)  <?mul(236,319){what() when()^'?who()mul$^mul(931,998)mul(251,79)mul(809,168)>[who()[)mul(734,171)#mul(879,259)select()*@when()mul(254,156)+mul(141,798)don't()mul(202,831)>>]/select()#!how()mul(601,846)^*,select(){from()mul(458,799)how(){why()'from()<mul(329,548)+ *^^mul(960,300)mul(615,752) &>who()?#how()~}<mul(916,199);mul(589,487)(/^)who()<@&what()mul(915,165$~?$select();~mul(83,883),select();;(where()]mul(839,114)[&how()-/-who() where()mul(594,959$+::$-what()select()mul(929,442)?why(){@mul(459,224)}}mul(699,260)!'+mul(911,579)^%(;<:&mul(644,689)how()where()how()mul(807,693)mul(400,649)select()^@what()mul(178,58)>where()[{who()select()do()from()why()from()mul(578,283)/*why(495,577)%select())'mul(475,993)@mul(673,70)mul(329,734@-how()when()*@where();&-[mul(703,672):+){who()<)&~)mul(770,646:)why())@~ @why()mul(867,88)mul(996,854)~what(804,700)where()mul(57,884);[how(),mul(593,934)what(324,653)when()}<!what(){/mul(824,86)^where(556,386)from()+mul(373,832)*'mul(729,891)!when()from()($#from()from()how()<mul(446,196)}%{?why()how()<mul(738,551)*what()'+mul(793,105){~^when()where()mul(325,94)@/>%mulselect()how()#/?*mul(588,116)*how()[don't()]@^:}how()??mul(881,189);+>@&<,why()#mul(168,770}mul(220,688)do()why()how();{- +(how()-mul(698,77)\n}where(174,722)({(/ *mul(582,635){what()~/:mul(470,940)where()'&mul(222,145)mul(880,163)$%*#>@how()how()%mul(25,277)(how(),-why(138,788)mul(41,986)]%when()from();<mul(362,999)select(),<^select()where()mul(768,107)select(){^}}(select()-mul(168,428)when()where()when()mul(179,874)when()-mul(835,227)]*%;how()'(mul(566,914) <do()@{ mul(438,926)[])!#;+mul(105,95)%[%~&?what(){mul(823,678)-what()mul(109,261)]{}mul(641,489)]~} +when()mul(921,460)who()(where():/ '+]mul(837,808)///?)$what()}mul(725,517])how()&[~[mul(837,920)how()~$&;</*mul(211,958))mul;(,where()mul(902,856)}!how()select()$:+mul(483,84)where()@>who(869,346)how()*why()([#mul(52,895)>don't()>+}!how()when()when()%#when()mul(765,168)from()mul(710,790)^#&~+select()mul(868,851)<where()mul(517,296)}how() +{]mul(401,32)mul(822,869)-when(940,602)who()~!where()mul(862,748)when()[:when()%why()mul(608,342)mul(357,681)%]$#!~)mul(303,13),how()what()mul(323,541)!()where()>from()how()->don't()^what()#>who()<;mul(722,400)~mul(217,112)from(90,168)!]when()$mul(806,665)%}>who()%-~mul(94,953)[-/*mul(322,64)when()mul(214,392)how()-{$:}don't() why(){,-who(703,968)!-mul(251,12)how()who()?from(538,173)&%@mul(318,562)how(),select()&,how()where()mul(239,777);~-mul(354,576)/~@$!^[when()mul(274,212)why()*mul(184,508)who()<@#do()]+where()mul(328,716)select();%mul(459,510)^'*when(632,422)+#where()how()select(260,585)^don't()^from()]$select()(who()how():%mul(35,792)mul(884,132)#mul(634,560),who()who();mul(759,868)/mul(688,411)why()mul(850,281)?--(>]mul(348,71)mul(725,441)who() select()*mul(373,244):why()who()<mul(554,7)from()^@&where()@^mul(561,205),why()}mul,>!,%&/select()mul(266,281):-what()',%select(769,776)[when()mul(478,447)when()^select()%mul(834,170)mul(867,978)}^-!how()how()@mul(884,586);!-}mul(681,431)who()[select(),}from()>@,)mul(358,409)+*/{when()@mul(798,937)^@mul(811,516)mul(683,569)when(222,647)what(240,706)why()mul(271,504),what()[)#mul(194,749)why())select()mul(452,63)[how()mul(675,660)<}*!-]~!'>mul(604,21)<where()mul(393,791)%who()/what()/~#mul(947,930)'-mul(361,577)>:who()?(>[mul(650,115)})]<;%mul(384,732)*mul(749,686)where()mul(524,177)^'+! >mul(698,944)*where()!^@]how()don't()-mul(573,409)$><&'*mul(985,814):from()mul(531,513)%]when() [!how()[]mul(322,404){;do()-mul(760,394)]mul(845,760)mul(691,945)why()}$/,mul(589,837)'$[[+&>mul(864,413)%~when()@'mul(650,851)$why()!where()!from()mul(437,509)why()mul(671,154)'+^#]how(),how()mul(248,241)mul(36,707):/~$%~>mul(415,475)select())(~(where())}mul(956,228)mul(310,43)how()$how()[how()~from()mul(230,10)what()'when()what()~why();:how()mul(219,566)how()*$#,>don't()[how()[,:]}%@%mul(907,64)@&why():]$)!mul(469,550))<]mul(328,818),from()why()]do()^'mul(525,979)how()<;]$+mul(451,350)~who()/'*<mul(824,446)^where()how()mul(212,756)</^mul(59,557)+mul(810,500)@/ -?:~+mul(957,32)how()}>what()]!$mul(322,822)mul(837,964),mul(997,114)??]how()mul(326,455)\n+!!who(111,197)<#mul(927,707)~@-()'mul(310,954)select()~;>%'don't()/mul(111,300)select()^#<,+-*when()where()mul(853,48)},where()<} do()>why()[* :mul(329,340)mul(775,432),what()]#select():mul[how()mul(975,305)don't() mul(380,23)<mul(610,328)$mul(547,22)how()where()select()'-when()select()how()*do()when()*&{@mul(509,794)'/^who(){what(){who()what()why(559,447)mul(527,574){*mul(640,883)~[, >^-^*mul(92,490)<?:#!);'mul(830,759)^-*~:?when()mul(658,459)who()do()%; %mul(901,84)@  {);mul(176,691)select()&]&[:[when()(who()mul(787,832)(({>!who()from()}@mul(482,720)/+mul(341,296)/%)%}$mul(199,41)?;select()mul(952,575)+%!;from()mul(629,288)from()mul(261,169) $when()mul(137,539)mul'who())who()# :,#!how()mul(712,429)^how()who()?;mul(572,59)>{ @mul(471,598)(>}select()*+mul(482,644)select()^why()select()+~!where(),mul(5,217)^>mul(927,157)mul(192,700)mul(598where()(}who()mul(914,246)mul(146,709),mul(545,12)mul(531,165)>#-&!*}mul(861,209)}{how()??(mul(469,555)/<:]}+%select()?mul(544,310){select()mul(549,681)+@^$ {?)%mul(625,262)%mul(658,715)-where()?/]who()* $&don't()what(366,880))mul(339,633)mul(282,510)-?'mul(538,886)mul(431,412)select()?# ]]mul(16,157)}}(!select()/mul(925,798);what(934,305)*[+-$why()what()mul(882,680)-/mul(554,323)/)where()when(),mul(384,806)~where()select()why()}what()[mul(291,811)&!how()&)where()@mul(675,825)when()mul(703,914)%-% ?@mul(733,721)/'*<mul(404,764)?!&mul(302,783)%/where();-,why()*mul(106,771)mul(422,483)!/?!from(746,984)don't()from(818,501)when()$;,'#>what()mul(452,139);>!how()(~how(681,256)&who(740,533)don't()#%why()+where()~mul(52,930):+ /do()&&:how()why()'mul(673,741)}where()don't()(who()mul(496,885)/mul(126,603)what()select()who(301,758)when()from()-mul(890,826)>]when()^>mul(231,474)!when(456,931)!where()[ how()+!-don't()mul(168,864)?mul(711,701) +^:: mul(243,505)!from()-+'?<select()mul(893,538)&-:mul(871,651)[from())(don't()({'@;&#mul(406,612)mul(52,182)[mul(785,121) mul(497,363)(who()(),mul(799,612))from()mul(518,886);&&%}who()^!mul(706,840)/mul(846,51)#>~;when()how():who()mul(482,341)<>!'where()why()}+mul(690,370)&#who()mul(682,266):@,+]^@-$mul(228,532)+#+where()from())'mul(547,133)@*mul(764,842)mul(181,500)*%!'/]mul(366,266)>](?>*<:why())mul(786,240)/ mul(980,844):@;![mul(932,294)[mul(582,252),when()?why(),~how()mul(619,83)why()??(^select()!mul(915,820)from()[+^don't();:# ]mul(421,350)'}how()mul(429,848)~$!select() ?/+/)mul(393,36)what(),%<-;+~mul(333,687)$(select(){{<select()what()mul(156,751) >-what(),mul(364,843))how()mul(148,833)how()'from()from():[}$mul(134,950)from()what() (what()mul(512,454)how(),mul(775,814):select(820,315),from()>who()how():**mul(329,285)where()?don't()!/->~/' mul(873,470)+-[(select()from()select()/why()?mul(911,768)what(805,778)mul(690,737)from()who())select()<~mul(248,530)mul(638,821)mul(218,217)(^why();&mul(684,550)"
	r, err := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	if err != nil {
		panic(err)
	}

	muls := r.FindAllString(input, -1)
	result := 0
	for _, mul := range muls {
		mul = strings.TrimPrefix(mul, "mul(")
		mul = strings.TrimSuffix(mul, ")")
		numbers := strings.Split(mul, ",")
		if len(numbers) != 2 {
			panic(fmt.Sprintf("numbers length is not 2. numbers: %v", numbers))
		}

		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		result += x * y
	}
	fmt.Println(muls)
	fmt.Println(result)
}
