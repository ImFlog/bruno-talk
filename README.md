# Let's not talk about Bruno

* [ ] UI prez a améliorer == c'est moche pour le moment
* [ ] Scénario de la démo à écrire quelque part et à répêter.
* [ ] Vérifier le timing !

## Scenario

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
10. Lancer la collection complète via CLI (penser à bien modifier la configuration de l'auth sur la route de création de produit)
   * `bru run bruno --env local`