<?php

function DBConnection (): PDO  {
    return new PDO('sqlite:./database/artbox.db');
}