# Реферат  

## по курсу "Логическое программирование"

### студент: Попов Матвей

### группа: М8О-208Б-20

## ТЕМА: Логические языки как путь к автоматическому решению задач компьютером

## Результат проверки

| Преподаватель     | Дата         |  Оценка       |
|-------------------|--------------|---------------|
| Сошников Д.В. |              |               |
| Левинская М.А.|              |               |

> *Комментарии проверяющих*

## Глава 1. Логическое программирование

### Введение в логическое программирование

Чаще всего, когда люди слышат о программировании или написании программ (чем, между прочим, программирование по своей сути и является), они представляют себе множество строк кода, в которых умелыми программистами ловко зашифрованы команды, выполняемые электронной вычислительной машиной (ЭВМ, по-нашему просто компьютер). Такого рода ассоциации у непосвящённых программистов возникают из-за огромного разнообразия *императивных* языков программирования. Как описывалось ранее, код, написанный на императивном языке программирования представляет собой набор команд, который примерно в одинаковой степени понимают и человек, и бездушная машина. Задача программиста — написать такую последовательность команд, чтобы компьютер правильно её понял, задача компьютера же заключается в том, чтобы безошибочно выполнить все эти команды и дать программисту результат, полученный в результате строгого следования инструкциям программиста. В девяноста девяти случаях из ста программист придерживается именно такой философии написания компьютерных программ (а лучше сказать *парадигме*), но программирование не такое уж и молодое ремесло, как может показаться на первый взгляд. За достаточно продолжительное время своего существования это ремесло успело обрасти множеством подходов к созданию программ. Такие подходы тоже называют словом «парадигма». Императивное программирование — самая старая и наиболее распространённая среди программистов любого ранга парадигма программирования, но сейчас мы почти не будем её затрагивать. Главным объектом моего реферата станет совсем иная парадигма (она же подход либо же философия написания программ) — логическое программирование.  

### Таки что же такое логическое программирование?

Поговорим ещё немного о разнообразии парадигм программирования. Я уже упоминал ранее, что программирование как род деятельности человека крайне разнообразно, и это разнообразие прежде всего отображается в разнообразии парадигм программирования. Во введении я постарался как можно более подробно рассказать об императивном программировании, теперь же настала пора рассказать о парадигме, которая наиболее сильно контрастирует с императивным программированием — *декларативной*. Если непосвящённый программист взглянет со стороны на работу «императивного» и «декларативного» программистов, он не заметит особой разницы. Действительно, оба программиста пишут программу, а компьютер её выполняет. Но программисты по-разному пишут программы, а компьютеры, очевидно, по-разному их выполняют, и такая разница не может быть объяснена разным мышлением программистов или разными составляющими вычислительных машин. «Декларативный» программист не пишет команды в своей программе, а его компьютер соответственно не выполняет их. Вместо этого программист даёт компьютеру нечто иное — некую модель ожидаемого результата выполнения программы, тем самым на плечи вычислительной машине ложится совсем иная задача — всеми правдами и неправдами найти все ответы, которые будут удовлетворять заданной человеком модели. Логическое программирование придерживается именно таким принципам, поэтому его относят к одному из многочисленных видов декларативного программирования. Основной отличительной особенностью логического программирования на фоне остального декларативного программирования является то, что в логическом программировании всё основано на законах математической логики (нетрудно догадаться, что именно отсюда и взялось название). В математической логике, как и в логическом программировании, всё держится на *утверждениях*. Утверждение может быть правильным либо неправильным. Правильность утверждения нужно либо доказать, либо опровергнуть, причём наше доказательство либо опровержение (тут уж как повезёт) должно быть основано на утверждениях, о которых мы заранее знаем, что они правильны (либо доказали правильность самостоятельно, либо они были даны изначально как правильные). В логическом программировании работу на определение правильности или неправильности утверждения берёт на себя компьютер, от человека требуется внести в программу набор правильных утверждений (тех самых, на которые будет опираться компьютер в своих доказательствах) и сформировать правила взаимодействия утверждений разного характера. В языке программирования Prolog утверждения представлены в форме *предикатов*. Помимо предикатов в Prolog есть правила, которые позволяют делать одни утверждения на основе других, тем самым обеспечивают взаимное действие различных предикатов. Таким образом, в логических языках программирования человек сделал попытку максимально приблизить работу компьютера к работе человеческого мозга. Действительно, ведь в повседневной жизни мы не думаем переменными, структурами данных или, не дай Бог, классами. Мы пользуемся немного более замысловатым инструментом — логикой, которая в свою очередь должна быть основана на фактах.  

### Откуда есть пошло логическое программирование?

Из предыдущего раздела можно сделать вывод, что логическое программирование является одной из ветвей декларативной парадигмы программирования. Если взять за истину то, что парадигма программирования зарождается одновременно с появлением первого языка программирования этой парадигмы, то логическое программирование вероятнее всего появилось на свет в одна тысяча девятьсот шестьдесят девятом году одновременно с языком программирования Planner. У этого языка были все черты, присущие остальным декларативным языкам программирования, а именно в нём не было переменных и оператора присваивания, но помимо этого Planner стал одним из пионеров техники *бэктрекинга*. Под этим словом понимают необходимость перебора всех возможных вариантов решения, чтобы найти правильные. Изучаемый мной язык программирования Prolog, появившийся на свет буквально на три года позже, а именно в одна тысяча девятьсот семьдесят втором году, сумел избавиться от столь нерационального метода нахождения ответов на поставленные вопросы. Вместо тупого перебора всех возможных вариантов Prolog строит из них дерево (подразумевается динамическая нелинейная структура данных), что сильно облегчает работу компьютеру. Prolog и по сей день остаётся самым распространённым языком логического программирования, к тому же этот язык оказал неоценимое влияние на Mercury — тоже довольно распространённый в наше время язык логического программирования, хотя конечно до популярности родителя ему ещё довольно далеко. Тем самым можно сделать вывод, что история развития логического программирования гораздо более скудна, чем, например у императивной ветви программирования. Это связано прежде всего с гораздо меньшей распространённостью логического программирования, но это не отменяет того факта, что логическое программирование всё-таки смогло найти прикладное применение, и, соответственно, человечество нуждается в том числе и в логическом программировании. Да, Prolog и Mercury не так распространены, как C++ и Python, но ведь не только в распространении измеряется значимость языка?

### Почему человечество нуждается в логическом программировании?

Окружающий наш мир крайне многогранен и из-за этого довольно сложно устроен. Тем не менее существует множество простых для понимания человеком правил, которым подчиняется абсолютно всё живое и неживое в этом мире (очень часто это происходит неосознанно). Одно из таких правил: то, без чего человек может спокойно обойтись, очень быстро им забывается. Ранее упоминалось, что логическое программирование появилось в шестидесятых годах прошлого века, при этом оно существует и по сей день и довольно активно используется. Из этого можно сделать логический (тонкая игра слов) вывод, что люди испытывают некую нужду в логическом программировании (либо просто пока не придумали ничего лучше, но это уже совсем другая история). Логические языки программирования могут оказаться весьма мощным инструментом, в первую очередь многие из них обладают полнотой по Тьюрингу (если очень коротко, с помощью логических языков программирования можно реализовать все те же алгоритмы, что и с помощью императивных языков программирования). В некоторых случаях решение задачи на языке Prolog может оказаться гораздо проще для понимания, чем на том же Python, однако стоит учитывать, что в таком случае программа на языке Prolog будет работать несколько медленнее. К задачам, традиционно решаемым на языке программирования Prolog, относится, например, анализ текста на естественном языке.

### Решение задач человеком и компьютером

Если мы говорим о решении логических задач, то человек с малых лет способен решать простейшие из них. Действительно, логика является одним из характерных свойств мышления человека, ведь ещё с младенчества вся информация об окружающем мире в человеческом мозге лучше всего закрепляется именно с помощью логических рассуждений. Ранее уже упоминалось, что логические языки программирования позволяют снабдить любую вычислительную машину таким мощным инструментом в решении задач, как логика, та самая логика, которой любой из нас обладает с рождения, причём для её использования не нужно многолетнее обучение различным законам и алгоритмам, как например происходит с математикой в школе. Значит ли это, что создание логических языков программирования есть попытка навязать бездушному компьютеру образ мышления человека? Отчасти да, потому что это самый очевидный способ научить компьютер решать хотя бы самые простые логические задачи. Тем самым человечество сделало самый первый и самый важный шаг к автоматическому решению задач компьютером.  

### Автоматическое решение заач компьютером

Ранее мы уже выяснили, что логическое программирование является отличным инструментом, который позволяет обучить компьютер человеческому образу мышления, но возможно ли таким образом добиться от вычислительной машины автоматического решения любой логической задачи, с которой может управиться человек? С помощью средств языков логического программирования мы можем построить модель задачи по её условию, тогда компьютер окажется вполне способен решать не только идентичную, но и схожие с таковой задачи, например, если в модификации задачи мы ограничимся изменением некоторых переменных. Но возможно ли создать такой алгоритм, который был бы способен на решение любой логической задачи? К сожалению, по описанному выше методу невозможно написать алгоритм решения логической задачи, не зная условия самой задачи, то есть у нас не получится написать такую программу, которая давала бы ответ на задачу, алгоритм решения которой не был обозначен в нашей программе. Вспоминается известная притча о том, что голодному человеку следует дать удочку, а не рыбу, так как с помощью удочки человек способен весьма продолжительное время кормить себя только что выловленной рыбой, а одной лишь рыбы без удочки человеку хватит лишь на один или два дня сытой жизни. Следовательно, компьютеру нужна удочка. Логическое программирование дало нам возможность использовать элементы человеческой логики при решении некоторых задач, возможно ли перенести человеческую логику в компьютер целиком?  

## Глава 2. Искусственный интеллект

### Введение в искусственный интеллект

В конце предыдущей главы мы практически вплотную подобрались к компьютеру, способному автоматически решать любые логические задачи. У человека способность решать логические (да и любые другие) задачи существует благодаря интеллекту, а это значит, что если мы хотим обучить компьютер решать задачи автоматически, то мы рано или поздно будем вынуждены снабдить его интеллектом. Такой интеллект будет называться искусственным, что означает воссозданный человеком по образцу природы. Появление искусственного интеллекта предсказывали практически в то же время, в какое зарождалось логическое программирование, а это может означать, что эти две вещи очень тесно связаны между собой. Несмотря на то, что термин «Искусственный интеллект» существует уже больше пятидесяти лет, до сих пор не до конца понятно, возможно ли создание «сильного» искусственного интеллекта, так как его создание подразумевает наделение компьютера способностью не только к решению задач, но и к творчеству.  

### «Сильный» искусственный интеллект

Очевидно, что основным инструментом человеческого интеллекта при решении различных задач является способность человека к мышлению. Несомненно, искусственный интеллект так же должен обладать способностью мыслить, чтобы научиться решать задачи. Способности мыслить и обучаться, принимать решения, воспринимать реальность являются ключевыми для «сильного» искусственного интеллекта. Под этим термином принято понимать такой искусственный интеллект, который способен решать задачи, то есть тот самый интеллект, который и является нашей целью. Уже довольно продолжительное время ведутся различные рассуждения о создании «сильного» искусственного интеллекта, а также о том, возможно ли его создание в принципе. На случай, если действительно появится что-либо похожее на «сильный» искусственный интеллект, существует несколько испытаний, с помощью которых можно подтвердить или опровергнуть способность машины мыслить, принимать решения или решать задачи.  

### Тест Тьюринга  

Идея Алана Тьюринга для определения способности к мышлению у машины возникла ещё до термина «искусственный интеллект». Сам тест Тьюринга заключается в том, что один из участников (будем называть его судья) проводит два диалога: между другим участником (будем называть его человек) и компьютером, способным поддерживать диалог на естественном языке. Судье неизвестно, кто из участников является компьютером, а кто является человеком. По окончании диалога судья должен определить, кто из участников кем является. Если судья не может ответить определённо, кто из участников кем является, считается, что машина прошла тест. При этом во время теста соблюдены следующие условия: участники, которые являются людьми, не видят друг друга; диалог ведётся в текстовом формате, чтобы проверить именно способность компьютера поддерживать диалог, а не распознавать либо синтезировать устную речь; между ответами всех участников проходят равные участки времени, чтобы невозможно было распознать компьютер по времени обработки сообщений. Несмотря на то, что идея этого теста появилась больше семидесяти лет назад, прохождение теста Тьюринга до сих пор является одной из обязательных способностей «сильного» искусственного интеллекта. Уже известны случаи прохождения теста Тьюринга, но обладают ли машины, которые прошли его, именно «сильным» искусственным интеллектом? Прежде чем ответить на этот вопрос, стоит разобраться в явлении китайской комнаты.  

### Китайская комната  

Данный мысленный эксперимент был предложен в одна тысяча девятьсот восьмидесятом году. Его принцип заключается в том, что существует изолированная комната, в которой находится первый участник эксперимента (будем называть его исполнитель). Исполнитель совершенно не знает китайского языка, однако в комнате помимо него есть книга, содержащая инструкции по составлению ответа на вопрос, заданный на китайском языке, однако в книге не содержится информация о значении иероглифов. Второй участник эксперимента (будем называть его наблюдатель) свободно владеет китайской письменностью, и может передавать в изолированную комнату различные вопросы на китайском языке, записанные на бумаге. Исполнитель должен точно следовать инструкциям из книги, чтобы подобрать правильные иероглифы для ответа, основываясь на иероглифах, которые содержатся в вопросе. Таким образом, исполнитель, который не владеет китайским языком и даже не обучается ему в процессе эксперимента, оказывается способен поддерживать диалог с наблюдателем исключительно на китайском языке. Если исполнителя в комнате заменить на компьютер, а книгу заменить на алгоритм, понятный машине, получим искусственный интеллект, который в перспективе способен пройти тест Тьюринга. Однако такой искусственный интеллект точно нельзя назвать «сильным», так как ответы компьютера основаны не на мыслительной деятельности, присущей «сильному» искусственному интеллекту, а на простом следовании алгоритму. В главе 1 уже упоминалось, что составление множества алгоритмов для решения многих задач не является верным путём, так как в процессе следования этим алгоритмам у компьютера не происходит процесс обучения. Если такому компьютеру дать задачу, для решения которой у него не будет алгоритма, он не сможет применить опыт решения прошлых задач, необходимый для её решения. Точно так же исполнитель не сможет составить вразумительный ответ на вопрос, для которого не будет инструкции в книге. Таким образом можно сделать вывод, что если некоторый компьютер прошёл тест Тьюринга, то это не означает, что этот компьютер обладает способность к обучению и принятию решений, а значит для автоматического решения задач он мало годится.  

### Что делать?

Значит ли то, что существование китайской комнаты в принципе отрицает возможность создания «сильного» искусственного интеллекта? Не стоит опускать руки. Китайская комната действительно не предполагает, что компьютер с заданным ограниченным и, что самое главное, нерасширяемым количеством инструкций после выхода из китайской комнаты будет способен решать любые задачи самостоятельно. Тем не менее, китайская комната дала нам одну очень важную подсказку. Ключом к способности машины автоматически решать любые задачи является способность её к обучению. Именно эта способность, характерная для человека, отделяет машину от автоматического решения задач.  

## Глава 3. Нейронная сеть

### Введение в нейронные сети

За свою историю, которая насчитывает более сорока тысяч лет, человек преодолел множество эпох: от первобытно-общинного строя до индустриального общества, а затем и до постиндустриального общества, которое сейчас иногда называют информационным обществом (неудивительно, что в некоторых кругах такого общества остро стоит вопрос искусственного интеллекта и компьютерного обучения). Эти периоды развития человечества можно характеризовать по-разному, например, уровнем развития культуры, традиций или права. Однако у различных народов были различные культуры и традиции (впрочем, довольно значимые различия в культурах разных народов сохранились и по сей день). Тем не менее существует один критерий, который позволяет характеризовать и отличать эпохи развития человеческого общества вне зависимости от различных стран и народов с достаточно высокой точностью. Этим самым критерием является технический прогресс, а именно технические достижения человечества. В период глобализации значимые научные открытия и изобретения практически моментально распространяются по всему миру, однако так было не всегда. Буквально несколько столетий назад, когда ещё не существовало столь развитой транспортной сети и уж тем более компьютеров и интернета, между двумя достаточно отдалёнными друг от друга государствами могла быть огромная пропасть в техническом развитии. Тем не менее именно технический прогресс наиболее точно отражает то, насколько развито человеческое общество. Человеческая история насчитывает немало гениальных изобретателей, привнесших то, что делает наш мир таким, каким мы с вами его знаем. Хороший изобретатель должен обладать немалой интуицией и огромным творческим потенциалом, а гениальный изобретатель помимо всего вышеперечисленного должен также обладать хорошим источником вдохновения. Никто не будет спорить, что в таком случае лучшим источником вдохновения для изобретателя является некто, кто в изобретательском ремесле стоит на качественно ином уровне. Этим изобретателем, который намного обогнал все человеческие умы, является природа, ведь именно природа, так как именно природа использует данные ей ресурсы максимально рационально, снабжая живые организмы наиболее эффективными инструментами для выживания. Действительно, много раз происходило такое, что человек просто подсматривал некоторые механизмы животных или растений и на их основе создавал похожие, а может быть даже и полностью копирующие природный оригинал, изобретение. Можно найти массу подобных примеров из различных эпох, одни оказывались удачным до такой степени, что в том или ином виде используются нами в повседневной жизни до сих пор, другие же ждала неудача, однако не будем вдаваться в столь тонкие подробности, чтобы не отходить далеко от поставленной темы. Перед программистами всего мира встала задача передать компьютеру снабдить искусственный интеллект тем, что отличает «сильный» искусственный интеллект от «слабого», а именно способность к обучению, подобной человеческой. Очевидно, что наиболее рациональным решением будет подсмотреть механизм мышления самого человека, так как мы уже знаем, что созданные природой механизмы являются наиболее эффективными в необходимой нам сфере деятельности (в нашем случае такой сферой является решение задач). Главным мыслительным инструментом (в том числе и обучаемым) является головной мозг. Человеческий мозг представляет собой огромное количество так называемых нейронов — клеток, способных обмениваться друг с другом информацией в виде электрических сигналов. Головной мозг человека (как и любого другого животного) разделяется на множество отделов, каждый из которых отвечает за выполнение различных функций, необходимых для поддержания жизнедеятельности организма. Помимо этого, существуют отделы, благодаря которым становится возможным обучение организма различным явлениям окружающей природы. Именно механизм нейронов является ключом к созданию обучаемого искусственного интеллекта. Создаваемые на основе нейронных связей головного мозга человека модели называются нейронными сетями.  

### Вывод

Нейронные сети качественно отличаются от традиционных алгоритмов именно способностью к обучению. Так как нейронные сети создаются именно по подобию человеческого мозга неудивительно, что их разработка уже приносит плоды и принесёт их ещё больше в совсем недалёком будущем. Уже сегодня нейронные сети способны читать текст, который написан на естественном языке, а также понимать человеческую речь; распознавать различные образы и объекты и даже составлять прогнозы в различных сферах жизни человека (например, финансовой). Это говорит нам о том, что нейронные сети обладают способностью сопоставлять входные данные и результат. Определение закономерностей является ключевой способностью, необходимой для обучения как человека, так и компьютера. Помимо этого, уже ведутся работы по обучению нейронных сетей искусству. Уже существуют нейронные сети, которые способны, например, самостоятельно создать картину в стиле того или иного художника, предварительно внимательно проанализировав реальные работы этого художника, а это нам говорит о том, что машина, обладающая нейронной сетью, наравне со способностью определять закономерности обладает также и воображением. Совсем скоро программисты смогут снабдить машину сознанием, тем самым будет создан искусственный интеллект, обладающий достаточным набором качеств, чтобы называться «сильным». Задачей логического программирования будет обучить такой интеллект навыкам решения задач с опорой на математическую логику, чтобы суждения и выводы машины подчинялись законам математической логики. Данный процесс необратим, совсем скоро машины не будут уступать в решении логических задач человеку, а через некоторое время и вовсе обгонят его во много раз, как это происходило, например, с вычислительными способностями человека и машины, а именно со скоростью вычислений.  

### Используемая литература

Зюзьков В.М. Математическое введение в декларативное программирование  
Карпов Ю.Г. Теория автоматов  
Девятков В.В. Системы искусственного интеллекта  
Горбань А.Н. Обучение нейронных сетей  
Поляков Г.И. О принципах нейронной организации мозга