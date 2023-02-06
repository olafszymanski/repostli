resource "random_password" "username" {
  length  = 16
  special = false
}

resource "random_password" "password" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

resource "aws_db_instance" "db" {
  db_name             = var.environment
  engine              = "postgres"
  engine_version      = "14.6"
  instance_class      = "db.t3.micro"
  allocated_storage   = 20
  username            = random_password.username.result
  password            = random_password.password.result
  skip_final_snapshot = false

  tags = {
    Name = "postgres-db-${var.environment}"
  }

  depends_on = [
    random_password.username,
    random_password.password
  ]
}

resource "aws_secretsmanager_secret" "db_credentials" {
  name = "db-credentials-${var.environment}"

  depends_on = [
    aws_db_instance.db
  ]
}

resource "aws_secretsmanager_secret_version" "db_credentials" {
  secret_id     = aws_secretsmanager_secret.db_credentials.id
  secret_string = <<EOF
  {
    "username": "${random_password.username.result}",
    "password": "${random_password.password.result}",
    "host": "${aws_db_instance.db.endpoint}",
    "port": ${aws_db_instance.db.port},
  }
EOF

  depends_on = [
    aws_secretsmanager_secret.db_credentials
  ]
}