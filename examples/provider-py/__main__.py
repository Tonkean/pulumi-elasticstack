import pulumi_elasticstack

pulumi_elasticstack.Provider(
    resource_name='test_Provider',
    elasticsearch=pulumi_elasticstack.ProviderElasticsearchArgs(
        username='USERNAME',
        password='PASSWORD',
        endpoints=['http://localhost:9200']
    )
)
