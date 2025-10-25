---
marp: true
theme: default
---

# Projet 2: The ArtBox

![bg right 80%](./img/logo.png)

## Une galerie d'art contemporain

### _OpenClassrooms_

_Par André M._

---

## 1. Résumé du projet

L'objectif de ce projet était de transformer un site web statique de galerie d'art en une application web dynamique.

**Livrables produits :**

- **Base de données SQLite :** Pour le stockage de données.
- **Affichage dynamique :** Les pages d'accueil et de détail des oeuvres affichent les données provenant de la base de données.
- **Formulaire d'ajout :** Un formulaire permettant d'ajouter de nouvelles oeuvres.
- **Validation côté serveur :** Nettoie les données soumises et gère les erreurs.

---

## 2. Étapes de réalisation

Le projet a été réalisé en plusieurs étapes clés, visibles à travers l'historique des commits :

1. [Initialisez le projet avec Git et GitHub](https://github.com/nutbreaker/projet_2/commit/bc26ca876c6cff217527813e559b7a2675c12880)
2. [Concevez la base de données du projet](https://github.com/nutbreaker/projet_2/commit/be8764ca4c862acf5822da46db6e5eb23c7cd7ce)
3. [Mettez à jour la page d'accueil du site](https://github.com/nutbreaker/projet_2/commit/7741765a9fe98869714d915d3feef9d0180d5863)
4. [Mettez à jour la page de détail d'une oeuvre](https://github.com/nutbreaker/projet_2/commit/435737a362b03726c7adbf30bd99f7eca18a8a92) et [refactor](https://github.com/nutbreaker/projet_2/commit/453a77e96a2bd6e84a95ce76b02edd9cdef56939)
5. [Validez le formulaire de création d’une oeuvre](https://github.com/nutbreaker/projet_2/commit/e07bec90710e40d74c7965081f749d40aea1af78)
6. [Insérez l'oeuvre en BDD](https://github.com/nutbreaker/projet_2/commit/d9c34064c4a6555f81621798a9b672804b350b02)

---

## 3. Ce que j'ai appris

- **Choix technique adapté au projet**, via l'utilisation de SQLite, qui privilégie la simplicité et la portabilité. La base de données est contenue dans un unique fichier, ce qui élimine le besoin d'un serveur dédié et simplifie considérablement le déploiement et la maintenance pour une application de cette envergure.
- **Naviguer dans une nouvelle code base**
- **Savoir quand s'écarter des consignes** afin de [simplifier](https://github.com/nutbreaker/projet_2/commit/dd5606f4042d7eb08c8bdc015eae4306f7166a8e) l'application

---

## 4. Points forts et réussites

- **Proposition de refactorisation**, utilisation de [`require_once`](https://github.com/nutbreaker/projet_2/commit/79d9de033cb0e4ed9df0c51ac20537fffea959ae) au lieu de `require`
- **Correction de [bugs](https://github.com/OpenClassrooms-Student-Center/PHP-P4-exercice/blob/96cd5cdaa65df86eaf6ccfad3d6935cacc2671a7/oeuvre.php#L7)** de l'application de base en utilisant la fonction [`exit`](https://github.com/nutbreaker/projet_2/commit/435737a362b03726c7adbf30bd99f7eca18a8a92#diff-285ea6d1c86e4e62fbda70daa1be3005754a3d5ca5efb745bb0dbc43402d64ddL22-L24)

```php
// Si l'URL ne contient pas d'id, on redirige sur la page d'accueil
if(empty($_GET['id'])) {
    header('Location: index.php');
}
```

---

## 5. Difficultés rencontrées

La principale difficulté a été de **ne faire que ce qui était demandé**. Lors de la réalisation d'un projet pour un client se tenir au cahier des charges permet de:

- Ne se concentrer que sur les exigences afin d'optimiser le temps
- Aller au-delà des attentes peut entraîner des coûts supplémentaires non prévus ou non facturés

---

## Conclusion

**Tout a une fin, sauf la banane qui en a deux.**
