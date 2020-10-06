'use strict';

module.exports = {
  up: (queryInterface, Sequelize) => {
    id: {
      allowNull: false,
          autoIncrement: true,
          primaryKey: true,
          type: Sequelize.INTEGER
    },
    name: {
      allowNull: true,
          unique: true,
          type: Sequelize.STRING
    },
    author_id: {
      allowNull: false,
          references: {
        model: {
          tableName: 'authors',
        },
        key: 'id'
      },
      type: Sequelize.INTEGER
    },
    created_at: {
      allowNull: false,
          type: Sequelize.DATE
    },
    updated_at: {
      allowNull: false,
          type: Sequelize.DATE
    }
  },

  down: (queryInterface, Sequelize) => {
    /*
      Add reverting commands here.
      Return a promise to correctly handle asynchronicity.

      Example:
      return queryInterface.dropTable('users');
    */
  }
};
