# Contribuer

Les contributions sont les bienvenues ! Voici comment participer.

## Signaler un problème

Ouvrez une [issue](../../issues) en décrivant le problème rencontré ou l'amélioration souhaitée.

## Proposer une modification

1. Forkez le dépôt
2. Créez une branche (`git checkout -b ma-modification`)
3. Effectuez vos changements
4. Vérifiez que la spec est valide :
   ```bash
   make lint
   ```
5. Si vous modifiez `openapi.yaml`, regénérez les clients :
   ```bash
   make generate-go
   ```
6. Committez et poussez votre branche
7. Ouvrez une Pull Request

## Conventions

- La documentation et les descriptions dans `openapi.yaml` sont rédigées en **français**
- Les noms de schémas et de champs suivent le vocabulaire métier FFBB
- Le code généré (`client.gen.go`) ne doit pas être modifié manuellement
