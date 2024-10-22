# We don't talk about Bruno

* [ ] Vérifier le timing !
* [ ] Si j'ai du temps, forker le projet de Alvaro et:
  * enlever les vagues
  * mettre le lien bluesky en bas
  * diminuer la taille de l'image de profile

## Intro (2min)

```bash
cd prez
yarn run dev
```

## Demo (10min)

Before running the scenario, make sure some products are inserted in the database.

1. `docker compose up -d`
2. Montrer l'UI de Bruno
3. Montrer la configuration d'une collection
   * Expliquer ce qu'est une collection => un ensemble de requêtes
   * Montrer qu'on peut avoir une configuration globale pour la collection (passer sur chaque menu)
4. Montrer la gestion des environnements
5. Montrer les routes non authentifiés
   * get all products -> montrer la notion de variables / secrets, les scripts, les assertions, les tests
   * get 1 product by id
6. Montrer la route de creation d'un produit -> fail car pas authentifié
7. Montrer la route de login et sa réponse
8. Montrer la configuration de l'auth au niveau de la collection
9. Activer l'heritage de la configuration de l'auth au niveau de la route de création de produit
10. Montrer rapidement l'arborescence des fichiers bruno dans mon IDE
11. Lancer la collection complète via CLI (penser à bien modifier la configuration de l'auth sur la route de création de produit)
   * `bru run bruno --env local`

## Comparison (3min)

* Reopen the slides