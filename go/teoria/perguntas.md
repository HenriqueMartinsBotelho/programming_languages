## Qual a diferença entre Go Routines e Threads

Embora go routines possam parecer threads tradicionais à primeira vista, elas são fundamentalmente diferentes em vários aspectos:

- **Leveza**: Go routines são muito mais leves do que threads. Enquanto uma thread pode exigir uma quantidade significativa de memória overhead por thread (tipicamente 1MB ou mais), uma go routine inicia com uma stack de apenas algumas kilobytes. Isso significa que é possível ter milhares, ou até mesmo milhões, de go routines em um programa Go sem esgotar a memória do sistema.

- **Gerenciamento**: O Go runtime gerencia as go routines internamente, incluindo a multiplexação das go routines em threads do sistema operacional. Isso simplifica o modelo de programação, já que o desenvolvedor não precisa se preocupar diretamente com a criação, a sincronização e o gerenciamento de threads.

- **Cooperatividade**: Enquanto threads tradicionais são preemptivas, go routines operam em um modelo mais cooperativo, com a troca de contexto ocorrendo durante operações de bloqueio ou quando explicitamente cedendo controle ao scheduler do Go usando `runtime.Gosched()`. Isso resulta em um uso mais eficiente dos recursos.
